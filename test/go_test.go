package main

import (
	connect_client ".."
	"context"
	b64 "encoding/base64"
	"fmt"
	ed "github.com/FactomProject/ed25519"
	factom "github.com/FactomProject/factom"
	"github.com/tjarratt/babble"
	"math/rand"
	"os"
	"strconv"
	"testing"
)

// Config struct definitions
type ApiConfig struct {
	AppId, AppKey, ChainId, Content string
}

type ChainsApiConfig struct {
	ExternalIds []string
}

type EntriesApiConfig struct {
	EntryHash   string
	ExternalIds []string
}

type IdentitiesApiConfig struct {
	IdChainId, EntryHash, SignerKey, SignerPrivateKey, OldKey, NewKey string
	Keys                                                              []string
}

type ProofsApiConfig struct {
	Height, Entry string
}

// Config initialization
var apiConfig = new(ApiConfig)
var chainsConfig = new(ChainsApiConfig)
var entriesConfig = new(EntriesApiConfig)
var identitiesConfig = new(IdentitiesApiConfig)
var proofsConfig = new(ProofsApiConfig)

var skipTests bool

func TestMain(m *testing.M) {
	fmt.Printf("Starting Golang test suite.\n")

	// Toggle selective test skip
	// Set to true to skip all tests
	// To run an individual test comment out the skip code at the top of the test case
	skipTests = false

	// Config assignment
	apiConfig = &ApiConfig{
		AppId:   "APP_ID",
		AppKey:  "APP_KEY",
		ChainId: "6fdeb23645cbfe1fe1b835eb31677662bcd0aec910f22d7f8c364e5b410295ca",
		Content: "VGVzdCBwb3N0IGNoYWluIC0gR29sYW5nIENsaWVudA==",
	}

	chainsConfig = &ChainsApiConfig{
		ExternalIds: []string{"VGVzdCBJZCAxIC0gR29sYW5nIENsaWVudA==", "VGVzdCBJZCAyIC0gR29sYW5nIENsaWVudA=="},
	}

	entriesConfig = &EntriesApiConfig{
		EntryHash:   "e689fea839469d6b841c38cdd666b702da4801debe798ac0b6f6d25dfaacb652",
		ExternalIds: []string{"ODQyYzczMDU4ZWJiOGJkMTVlNWU2NGY2ZGUxOGYyNDdlZmYwMWMyMjYxZDI0YjcwYTdmZWY4YmY5Mjg3ODA0Yg=="},
	}

	identitiesConfig = &IdentitiesApiConfig{
		IdChainId:        "842c73058ebb8bd15e5e64f6de18f247eff01c2261d24b70a7fef8bf9287804b",
		EntryHash:        "idpub1zU3PRb6krwYHQbCyT8pLnXP1TztmvmczmYFus3BHgQYB7KdDs",
		SignerKey:        "idpub2XWG5Je7BKVM7qVCyWv1MkWmzdQPTg1NCBbbh3WNBb9EkrbtDp",
		SignerPrivateKey: "idsec1QdRgKM3Xmx2pXhJJAdRdev9Ch6TTBZV6FUCN23LoenyZbGf3t",
		OldKey:           "idpub3bLQKi7E8py7REgqdYbtMEoZaoqAhgDFphoHJmZ2qQcFiosZyj",
		NewKey:           "idpub3XaYqL6yqEBHBUMPPTgjYGnNU2gZP2t5iaPB829XzRSrZ1ao9k",
		Keys:             []string{"idpub2ZsMuKkkfhZRzPe8xij8a1CZ2vcDH5tNagnSPkRGPAEuLGUtJJ", "idpub3SCdbbNebT14an2mh8mDVXRWsiZRaZnXccdupEDrcrUP3PJBDT"},
	}

	proofsConfig = &ProofsApiConfig{
		Height: "22",
		Entry:  "287f8b117c508e0596e23109a56b08a2040c695b19ba966a4e996b0da3633fd0",
	}

	exitVal := m.Run()
	os.Exit(exitVal)
}

// Chains API

func TestGetChainByID(t *testing.T) {

	if skipTests {
		// Comment out the line below with skipTests in TestMain() set to true to run this test individually
		t.Skip()
	}

	client := connect_client.NewAPIClient(connect_client.NewConfiguration())

	keys := connect_client.APIKey{AppId: apiConfig.AppId, AppKey: apiConfig.AppKey}

	auth := context.WithValue(context.Background(), connect_client.ContextAPIKey, keys)

	_, res, err := client.ChainsApi.GetChainByID(auth, apiConfig.ChainId)

	if err != nil {
		t.Errorf("Error occured: %v \n\n", err)
		t.Errorf("Detailed response: %v", res)
	}
}
func TestGetChains(t *testing.T) {

	if skipTests {
		// Comment out the line below with skipTests in TestMain() set to true to run this test individually
		t.Skip()
	}

	client := connect_client.NewAPIClient(connect_client.NewConfiguration())

	keys := connect_client.APIKey{AppId: apiConfig.AppId, AppKey: apiConfig.AppKey}

	auth := context.WithValue(context.Background(), connect_client.ContextAPIKey, keys)

	_, res, err := client.ChainsApi.GetChains(auth, nil)

	if err != nil {
		t.Errorf("Error occured: %v \n\n", err)
		t.Errorf("Detailed response: %v", res)
	}
}

func TestPostChain(t *testing.T) {

	if skipTests {
		// Comment out the line below with skipTests in TestMain() set to true to run this test individually
		t.Skip()
	}

	client := connect_client.NewAPIClient(connect_client.NewConfiguration())

	keys := connect_client.APIKey{AppId: apiConfig.AppId, AppKey: apiConfig.AppKey}

	auth := context.WithValue(context.Background(), connect_client.ContextAPIKey, keys)

	chainCreate := connect_client.ChainCreate{ExternalIds: chainsConfig.ExternalIds,
		Content: apiConfig.Content}

	_, res, err := client.ChainsApi.PostChain(auth, chainCreate)

	if err != nil {
		t.Errorf("Error occured: %v \n\n", err)
		t.Errorf("Detailed response: %v", res)
	}
}

func TestPostChainSearch(t *testing.T) {

	if skipTests {
		// Comment out the line below with skipTests in TestMain() set to true to run this test individually
		t.Skip()
	}

	client := connect_client.NewAPIClient(connect_client.NewConfiguration())

	keys := connect_client.APIKey{AppId: apiConfig.AppId, AppKey: apiConfig.AppKey}

	auth := context.WithValue(context.Background(), connect_client.ContextAPIKey, keys)

	searchBody := connect_client.SearchBody{ExternalIds: chainsConfig.ExternalIds}

	_, res, err := client.ChainsApi.PostChainSearch(auth, searchBody, nil)

	if err != nil {
		t.Errorf("Error occured: %v \n\n", err)
		t.Errorf("Detailed response: %v", res)
	}
}

// Entries API

func TestGetEntriesByChainID(t *testing.T) {

	if skipTests {
		// Comment out the line below with skipTests in TestMain() set to true to run this test individually
		t.Skip()
	}

	client := connect_client.NewAPIClient(connect_client.NewConfiguration())

	keys := connect_client.APIKey{AppId: apiConfig.AppId, AppKey: apiConfig.AppKey}

	auth := context.WithValue(context.Background(), connect_client.ContextAPIKey, keys)

	_, res, err := client.EntriesApi.GetEntriesByChainID(auth, apiConfig.ChainId, nil)

	if err != nil {
		t.Errorf("Error occured: %v \n\n", err)
		t.Errorf("Detailed response: %v", res)
	}
}

func TestGetEntryByHash(t *testing.T) {

	if skipTests {
		// Comment out the line below with skipTests in TestMain() set to true to run this test individually
		t.Skip()
	}

	client := connect_client.NewAPIClient(connect_client.NewConfiguration())

	keys := connect_client.APIKey{AppId: apiConfig.AppId, AppKey: apiConfig.AppKey}

	auth := context.WithValue(context.Background(), connect_client.ContextAPIKey, keys)

	_, res, err := client.EntriesApi.GetEntryByHash(auth, apiConfig.ChainId,
		entriesConfig.EntryHash)

	if err != nil {
		t.Errorf("Error occured: %v \n\n", err)
		t.Errorf("Detailed response: %v", res)
	}
}

func TestGetFirstEntry(t *testing.T) {

	if skipTests {
		// Comment out the line below with skipTests in TestMain() set to true to run this test individually
		t.Skip()
	}

	client := connect_client.NewAPIClient(connect_client.NewConfiguration())

	keys := connect_client.APIKey{AppId: apiConfig.AppId, AppKey: apiConfig.AppKey}

	auth := context.WithValue(context.Background(), connect_client.ContextAPIKey, keys)

	_, res, err := client.EntriesApi.GetFirstEntry(auth, apiConfig.ChainId)

	if err != nil {
		t.Errorf("Error occured: %v \n\n", err)
		t.Errorf("Detailed response: %v", res)
	}
}

func TestGetLastEntry(t *testing.T) {

	if skipTests {
		// Comment out the line below with skipTests in TestMain() set to true to run this test individually
		t.Skip()
	}

	client := connect_client.NewAPIClient(connect_client.NewConfiguration())

	keys := connect_client.APIKey{AppId: apiConfig.AppId, AppKey: apiConfig.AppKey}

	auth := context.WithValue(context.Background(), connect_client.ContextAPIKey, keys)

	_, res, err := client.EntriesApi.GetLastEntry(auth, apiConfig.ChainId)

	if err != nil {
		t.Errorf("Error occured: %v \n\n", err)
		t.Errorf("Detailed response: %v", res)
	}
}

func TestPostEntriesSearch(t *testing.T) {

	if skipTests {
		// Comment out the line below with skipTests in TestMain() set to true to run this test individually
		t.Skip()
	}

	client := connect_client.NewAPIClient(connect_client.NewConfiguration())

	keys := connect_client.APIKey{AppId: apiConfig.AppId, AppKey: apiConfig.AppKey}

	auth := context.WithValue(context.Background(), connect_client.ContextAPIKey, keys)

	searchBody := connect_client.SearchBody{ExternalIds: entriesConfig.ExternalIds}

	_, res, err := client.EntriesApi.PostEntriesSearch(auth, apiConfig.ChainId, searchBody, nil)

	if err != nil {
		t.Errorf("Error occured: %v \n\n", err)
		t.Errorf("Detailed response: %v", res)
	}
}

func TestPostEntryToChainID(t *testing.T) {

	if skipTests {
		// Comment out the line below with skipTests in TestMain() set to true to run this test individually
		t.Skip()
	}

	client := connect_client.NewAPIClient(connect_client.NewConfiguration())

	keys := connect_client.APIKey{AppId: apiConfig.AppId, AppKey: apiConfig.AppKey}

	auth := context.WithValue(context.Background(), connect_client.ContextAPIKey, keys)

	entryCreate := connect_client.EntryCreate{ExternalIds: entriesConfig.ExternalIds,
		Content: apiConfig.Content}

	_, res, err := client.EntriesApi.PostEntryToChainID(auth, apiConfig.ChainId, entryCreate)

	if err != nil {
		t.Errorf("Error occured: %v \n\n", err)
		t.Errorf("Detailed response: %v", res)
	}
}

// Identities API

func TestGetIdChainbyId(t *testing.T) {

	if skipTests {
		// Comment out the line below with skipTests in TestMain() set to true to run this test individually
		t.Skip()
	}

	client := connect_client.NewAPIClient(connect_client.NewConfiguration())

	keys := connect_client.APIKey{AppId: apiConfig.AppId, AppKey: apiConfig.AppKey}

	auth := context.WithValue(context.Background(), connect_client.ContextAPIKey, keys)

	_, res, err := client.IdentitiesApi.GetIdChainbyId(auth, identitiesConfig.IdChainId)

	if err != nil {
		t.Errorf("Error occured: %v \n\n", err)
		t.Errorf("Detailed response: %v", res)
	}
}

func TestGetKeybyEntryHash(t *testing.T) {

	if skipTests {
		// Comment out the line below with skipTests in TestMain() set to true to run this test individually
		t.Skip()
	}

	client := connect_client.NewAPIClient(connect_client.NewConfiguration())

	keys := connect_client.APIKey{AppId: apiConfig.AppId, AppKey: apiConfig.AppKey}

	auth := context.WithValue(context.Background(), connect_client.ContextAPIKey, keys)

	_, res, err := client.IdentitiesApi.GetKeybyEntryHash(auth, identitiesConfig.IdChainId, identitiesConfig.EntryHash)

	if err != nil {
		t.Errorf("Error occured: %v \n\n", err)
		t.Errorf("Detailed response: %v", res)
	}
}

func TestGetKeysbyIdChainId(t *testing.T) {

	if skipTests {
		// Comment out the line below with skipTests in TestMain() set to true to run this test individually
		t.Skip()
	}

	client := connect_client.NewAPIClient(connect_client.NewConfiguration())

	keys := connect_client.APIKey{AppId: apiConfig.AppId, AppKey: apiConfig.AppKey}

	auth := context.WithValue(context.Background(), connect_client.ContextAPIKey, keys)

	_, res, err := client.IdentitiesApi.GetKeysbyIdChainId(auth, identitiesConfig.IdChainId, nil)

	if err != nil {
		t.Errorf("Error occured: %v \n\n", err)
		t.Errorf("Detailed response: %v", res)
	}
}

func TestPostIdChain(t *testing.T) {

	if skipTests {
		// Comment out the line below with skipTests in TestMain() set to true to run this test individually
		t.Skip()
	}

	client := connect_client.NewAPIClient(connect_client.NewConfiguration())

	keys := connect_client.APIKey{AppId: apiConfig.AppId, AppKey: apiConfig.AppKey}

	auth := context.WithValue(context.Background(), connect_client.ContextAPIKey, keys)

	babble := babble.NewBabbler()

	// Generate a unique identity name of the form Golang-Client-Test-<randomWord>-<randomWord>-<randomNumber> to ensure uniqueness
	name := "Golang-Client-Test" + "-" + babble.Babble() + "-" + strconv.Itoa(rand.Intn(100))

	encodedName := b64.StdEncoding.EncodeToString([]byte(name))

	identityCreate := connect_client.IdentityCreate{Names: []string{string(encodedName)},
		Keys: identitiesConfig.Keys}

	_, res, err := client.IdentitiesApi.PostIdChain(auth, identityCreate)

	if err != nil {
		t.Errorf("Error occured: %v \n\n", err)
		t.Errorf("Detailed response: %v", res)
	}
}

func TestPostKeytoIdChainId(t *testing.T) {

	t.Skip("Skipping TestPostKeytoIdChainId until signature generation is implemented.")

	if skipTests {
		// Comment out the line below with skipTests in TestMain() set to true to run this test individually
		t.Skip()
	}

	client := connect_client.NewAPIClient(connect_client.NewConfiguration())

	keys := connect_client.APIKey{AppId: "4823d718", AppKey: "99746352cdaf86bec57ae23544fc6ee4"}

	auth := context.WithValue(context.Background(), connect_client.ContextAPIKey, keys)

	babble := babble.NewBabbler()

	// Generate a unique identity name of the form Golang-Client-Test-<randomWord>-<randomWord>-<randomNumber> to ensure uniqueness
	name := "Golang-Client-Test" + "-" + babble.Babble() + "-" + strconv.Itoa(rand.Intn(100))

	encodedName := b64.StdEncoding.EncodeToString([]byte(name))

	factomIdentityKeys := factom.NewIdentityKey()

	identityCreate := connect_client.IdentityCreate{Names: []string{encodedName},
		Keys: []string{factomIdentityKeys.PubString(), identitiesConfig.OldKey}}

	postedChainData, res, err := client.IdentitiesApi.PostIdChain(auth, identityCreate)

	if err != nil {
		t.Errorf("Error occured: %v \n\n", err)
		t.Errorf("Detailed response: %v", res)
	}

	createdChainID := postedChainData.ChainId

	signature := ed.Sign(factomIdentityKeys.SecFixed(), []byte(createdChainID+identitiesConfig.OldKey+identitiesConfig.NewKey))

	identityKeyReplace := connect_client.IdentityKeyReplace{OldKey: identitiesConfig.OldKey,
		NewKey: identitiesConfig.NewKey, SignerKey: factomIdentityKeys.PubString(), Signature: string(signature[:])}

	_, res, err = client.IdentitiesApi.PostKeytoIdChainId(auth, createdChainID, identityKeyReplace)

	if err != nil {
		t.Errorf("Error occured: %v \n\n", err)
		t.Errorf("Detailed response: %v", res)
	}
}

// Info Api
func TestGetApiInfo(t *testing.T) {

	if skipTests {
		// Comment out the line below with skipTests in TestMain() set to true to run this test individually
		t.Skip()
	}

	client := connect_client.NewAPIClient(connect_client.NewConfiguration())

	keys := connect_client.APIKey{AppId: apiConfig.AppId, AppKey: apiConfig.AppKey}

	auth := context.WithValue(context.Background(), connect_client.ContextAPIKey, keys)

	_, res, err := client.InfoApi.GetApiInfo(auth)

	if err != nil {
		t.Errorf("Error occured: %v \n\n", err)
		t.Errorf("Detailed response: %v", res)
	}
}

// Proofs Api

func TestGetAnchorbySearch(t *testing.T) {

	if skipTests {
		// Comment out the line below with skipTests in TestMain() set to true to run this test individually
		t.Skip()
	}

	client := connect_client.NewAPIClient(connect_client.NewConfiguration())

	keys := connect_client.APIKey{AppId: apiConfig.AppId, AppKey: apiConfig.AppKey}

	auth := context.WithValue(context.Background(), connect_client.ContextAPIKey, keys)

	_, res, err := client.ProofsApi.GetAnchorbySearch(auth, proofsConfig.Height)

	if err != nil {
		t.Errorf("Error occured: %v \n\n", err)
		t.Errorf("Detailed response: %v", res)
	}
}

func TestGetReceiptsbyEntry(t *testing.T) {

	if skipTests {
		// Comment out the line below with skipTests in TestMain() set to true to run this test individually
		t.Skip()
	}

	client := connect_client.NewAPIClient(connect_client.NewConfiguration())

	keys := connect_client.APIKey{AppId: apiConfig.AppId, AppKey: apiConfig.AppKey}

	auth := context.WithValue(context.Background(), connect_client.ContextAPIKey, keys)

	_, res, err := client.ProofsApi.GetReceiptsbyEntry(auth, proofsConfig.Entry)

	if err != nil {
		t.Errorf("Error occured: %v \n\n", err)
		t.Errorf("Detailed response: %v", res)
	}
}

func TestGetSearch(t *testing.T) {

	if skipTests {
		// Comment out the line below with skipTests in TestMain() set to true to run this test individually
		t.Skip()
	}

	client := connect_client.NewAPIClient(connect_client.NewConfiguration())

	keys := connect_client.APIKey{AppId: apiConfig.AppId, AppKey: apiConfig.AppKey}

	auth := context.WithValue(context.Background(), connect_client.ContextAPIKey, keys)

	getSearchOpts := connect_client.GetSearchOpts{Term: proofsConfig.Entry}

	_, res, err := client.ProofsApi.GetSearch(auth, &getSearchOpts)

	if err != nil {
		t.Errorf("Error occured: %v \n\n", err)
		t.Errorf("Detailed response: %v", res)
	}
}
