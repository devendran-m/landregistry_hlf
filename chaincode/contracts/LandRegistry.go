package contracts

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

//SmartContract that provides functions for managing a land registration
type LandRegistry struct {
	contractapi.Contract
}

//LandRec describes basic details of the Land/Property
type PropRec struct {
	PropType  string `json:"proptype"`
	PropCity  string `json:"propcity"`
	PropState string `json:"propstate"`
	PropSqFt  string `json:"propsqft"`
	PropOwner string `json:"propowner"`
}

//Query Land/Property details
type QueryResult struct {
	Key    string `json:"Key"`
	Record *PropRec
}

//InitLedger - initialize with a set of properties
func (s *LandRegistry) InitLedger(ctx contractapi.TransactionContextInterface) error {
	proprecs := []PropRec{
		PropRec{PropType: "Flat", PropCity: "Chennai", PropState: "TN", PropSqFt: "1200", PropOwner: "Dev"},
		PropRec{PropType: "Ind House", PropCity: "Bengaluru", PropState: "KA", PropSqFt: "3200", PropOwner: "Abraham"},
		PropRec{PropType: "Res Plot", PropCity: "Coimbatore", PropState: "TN", PropSqFt: "4000", PropOwner: "Jagan"},
		PropRec{PropType: "Res Villa", PropCity: "Palakkad", PropState: "KL", PropSqFt: "4800", PropOwner: "John"},
		PropRec{PropType: "Farm Land", PropCity: "Coimbatore", PropState: "TN", PropSqFt: "100000", PropOwner: "Fasil"},
		PropRec{PropType: "Commercial Bldg", PropCity: "Chennai", PropState: "TN", PropSqFt: "1200", PropOwner: "Hema"},
	}
	for i, proprec := range proprecs {
		propAsBytes, _ := json.Marshal(proprec)
		err := ctx.GetStub().PutState("PROP"+strconv.Itoa(i), propAsBytes)

		if err != nil {
			return fmt.Errorf("failed to put to world state. %s", err.Error())
		}
	}
	return nil
}

//Register a Property on to the ledger
//Prop Counter or Transaction Hash could be used to assign Prop ID. Prod grade approach wud be generating in middleware - feedback
func (s *LandRegistry) CreateProp(ctx contractapi.TransactionContextInterface, propID string, propType string, propCity string, propSt string, propSqFt string, propOwn string) error {
	proprec := PropRec{
		PropType:  propType,
		PropCity:  propCity,
		PropState: propSt,
		PropSqFt:  propSqFt,
		PropOwner: propOwn,
	}

	propAsBytes, _ := json.Marshal(proprec)

	return ctx.GetStub().PutState(propID, propAsBytes)
}

//Query a Property based on Property ID

func (s *LandRegistry) QueryProp(ctx contractapi.TransactionContextInterface, propID string) (*PropRec, error) {
	propAsBytes, err := ctx.GetStub().GetState(propID)

	if err != nil {
		return nil, fmt.Errorf("failed to read from world state. %s", err.Error())
	}

	if propAsBytes == nil {
		return nil, fmt.Errorf("%s does not exist", propID)
	}

	proprec := new(PropRec)
	_ = json.Unmarshal(propAsBytes, proprec)

	return proprec, nil
}

// QueryAllCars returns all cars found in world state
//Feedback - When generating IDs in a dynamic way, this st key and end key option may not work
func (s *LandRegistry) ListAllProps(ctx contractapi.TransactionContextInterface) ([]QueryResult, error) {
	startKey := ""
	endKey := ""

	resultsIterator, err := ctx.GetStub().GetStateByRange(startKey, endKey)

	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	results := []QueryResult{}

	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()

		if err != nil {
			return nil, err
		}

		prop := new(PropRec)
		_ = json.Unmarshal(queryResponse.Value, prop)

		queryResult := QueryResult{Key: queryResponse.Key, Record: prop}
		results = append(results, queryResult)
	}

	return results, nil
}

//Property Ownership Transfer using Property ID
func (s *LandRegistry) ChangePropOwner(ctx contractapi.TransactionContextInterface, propID string, newPropOwner string) error {
	proprec, err := s.QueryProp(ctx, propID)

	if err != nil {
		return err
	}

	proprec.PropOwner = newPropOwner

	propAsBytes, _ := json.Marshal(proprec)

	return ctx.GetStub().PutState(propID, propAsBytes)
}

type PropOwnerRep struct {
	PropType string `json:"proptype"`
}

//Rich Queries Here
//GetAllPropsforOwner : Get the list of properites for the given owner
func (s *LandRegistry) GetAllPropsforOwner(ctx contractapi.TransactionContextInterface, powner string) ([]*PropOwnerRep, error) {
	//queryString := fmt.Sprintf(`{"selector":{"docType":"PropRec","powner": "%s"}}`, powner)
	queryString := fmt.Sprintf(`{"selector":{"propowner": "%s"}}`, powner)

	resultsIterator, err := ctx.GetStub().GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var report []*PropOwnerRep
	for resultsIterator.HasNext() {
		queryResult, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		var propOwnRep PropOwnerRep
		err = json.Unmarshal(queryResult.Value, &propOwnRep)
		if err != nil {
			return nil, err
		}
		report = append(report, &propOwnRep)

		fmt.Println("Report", report)
	}
	return report, nil
}
