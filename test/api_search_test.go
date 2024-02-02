/*
Manticore Search Client

Testing SearchAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
	openapiclient "github.com/manticoresoftware/manticoresearch-go"
)

func Test_openapi_SearchAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	configuration.Servers[0].URL = "http://localhost:9408"
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SearchAPIService Percolate", func(t *testing.T) {

		apiClient.UtilsAPI.Sql(context.Background()).Body("DROP TABLE IF EXISTS products").Execute()
        apiClient.UtilsAPI.Sql(context.Background()).Body("CREATE TABLE IF NOT EXISTS products(title text, color string) type='pq'").Execute()
        
        indexDoc := map[string]interface{} {"query": "@title shoes", "filters": "color='red'"}
    	indexReq := openapiclient.NewInsertDocumentRequest("products", indexDoc)
    	apiClient.IndexAPI.Insert(context.Background()).InsertDocumentRequest(*indexReq).Execute()
    	
    	indexDoc = map[string]interface{} {"query": "@title boots", "filters": "color IN ('blue', 'green')"}
    	indexReq = openapiclient.NewInsertDocumentRequest("products", indexDoc)
    	apiClient.IndexAPI.Insert(context.Background()).InsertDocumentRequest(*indexReq).Execute()

    	indexReq = openapiclient.NewInsertDocumentRequest("products", indexDoc)
    	indexReq.SetId(2)
    	apiClient.IndexAPI.Insert(context.Background()).InsertDocumentRequest(*indexReq).Execute()

		query := map[string]interface{} {"document": map[string]interface{} {}}
		percQuery := openapiclient.NewPercolateRequestQuery(query)
		percReq := openapiclient.NewPercolateRequest(*percQuery)
		resp, httpRes, err := apiClient.SearchAPI.Percolate(context.Background(), "products").PercolateRequest(*percReq).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
		
		outRes, outErr := json.Marshal(resp)
    	require.Nil(t, outErr)
    	fmt.Printf("%+v\n", string(outRes[:]))
    	
    	doc := map[string]interface{} {"title": "nice pair of boots", "color": "blue"}
    	docs := []map[string]interface{} {doc}
    	query = map[string]interface{} {"documents": docs}

    	percQuery = openapiclient.NewPercolateRequestQuery(query)
		percReq = openapiclient.NewPercolateRequest(*percQuery)
		resp, httpRes, err = apiClient.SearchAPI.Percolate(context.Background(), "products").PercolateRequest(*percReq).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
		
		outRes, outErr = json.Marshal(resp)
    	require.Nil(t, outErr)
    	fmt.Printf("%+v\n", string(outRes[:]))

	})

	t.Run("Test SearchAPIService Search", func(t *testing.T) {

		testBasicSearch := func (search openapiclient.ApiSearchRequest, searchReq *openapiclient.SearchRequest) {
			query := map[string]interface{} {"query_string": "Star"};
			searchReq.SetQuery(query);
			searchReq.SetLimit(10);
			searchReq.SetTrackScores(false);
			resp, httpRes, err := search.SearchRequest(*searchReq).Execute()
			fmt.Printf("%+v\n", httpRes)
				
			require.Nil(t, err)
			require.NotNil(t, resp)
			assert.Equal(t, 200, httpRes.StatusCode)
			
			outRes, outErr := json.Marshal(resp)
    		require.Nil(t, outErr)
    		fmt.Printf("%+v\n", string(outRes[:]))
    		
    		options := map[string]interface{} {"cutoff": 2, "ranker": "bm25"}
			searchReq.SetOptions(options)

			resp, httpRes, err = search.SearchRequest(*searchReq).Execute()
				
			require.Nil(t, err)
			require.NotNil(t, resp)
			assert.Equal(t, 200, httpRes.StatusCode)
			
			outRes, outErr = json.Marshal(resp)
    		require.Nil(t, outErr)
    		fmt.Printf("%+v\n", string(outRes[:]))
    		
    		includes := []string {"title", "year", "rating"}
    		excludes := []string {"code"}
			source := map[string]interface{} {"includes": includes, "excludes": excludes}
			searchReq.SetSource(source)		
			
			resp, httpRes, err = search.SearchRequest(*searchReq).Execute()
				
			require.Nil(t, err)
			require.NotNil(t, resp)
			assert.Equal(t, 200, httpRes.StatusCode)
			
			outRes, outErr = json.Marshal(resp)
    		require.Nil(t, outErr)
    		fmt.Printf("%+v\n", string(outRes[:]))			
		}
		
		apiClient.UtilsAPI.Sql(context.Background()).Body("DROP TABLE IF EXISTS movies").Execute()
		resp, httpRes, err := apiClient.UtilsAPI.Sql(context.Background()).Body("CREATE TABLE movies (title text, plot text, year integer, rating float, code multi)").Execute()
    	fmt.Printf("test %+v\n", httpRes)
    	require.Nil(t, err)
		require.NotNil(t, resp)
		
    
		docs := [4]string {
			"{\"insert\": {\"index\" : \"movies\", \"id\" : 1, \"doc\" : {\"title\" : \"Star Trek 2: Nemesis\", \"plot\": \"The Enterprise is diverted to the Romulan homeworld Romulus, supposedly because they want to negotiate a peace treaty. Captain Picard and his crew discover a serious threat to the Federation once Praetor Shinzon plans to attack Earth.\", \"year\": 2002, \"rating\": 6.4, \"code\": [1,2,3]}}}",
	        "{\"insert\": {\"index\" : \"movies\", \"id\" : 2, \"doc\" : {\"title\" : \"Star Trek 1: Nemesis\", \"plot\": \"The Enterprise is diverted to the Romulan homeworld Romulus, supposedly because they want to negotiate a peace treaty. Captain Picard and his crew discover a serious threat to the Federation once Praetor Shinzon plans to attack Earth.\", \"year\": 2001, \"rating\": 6.5, \"code\": [1,12,3]}}}",
	        "{\"insert\": {\"index\" : \"movies\", \"id\" : 3, \"doc\" : {\"title\" : \"Star Trek 3: Nemesis\", \"plot\": \"The Enterprise is diverted to the Romulan homeworld Romulus, supposedly because they want to negotiate a peace treaty. Captain Picard and his crew discover a serious threat to the Federation once Praetor Shinzon plans to attack Earth.\", \"year\": 2003, \"rating\": 6.6, \"code\": [11,2,3]}}}",
	        "{\"insert\": {\"index\" : \"movies\", \"id\" : 4, \"doc\" : {\"title\" : \"Star Trek 4: Nemesis\", \"plot\": \"The Enterprise is diverted to the Romulan homeworld Romulus, supposedly because they want to negotiate a peace treaty. Captain Picard and his crew discover a serious threat to the Federation once Praetor Shinzon plans to attack Earth.\", \"year\": 2003, \"rating\": 6, \"code\": [1,2,4]}}}",					        	
		}
   		bulkResp, httpRes, err := apiClient.IndexAPI.Bulk(context.Background()).Body(strings.Join(docs[:], "\n")).Execute()
   		fmt.Printf("test %+v\n", httpRes)
    	require.Nil(t, err)
		require.NotNil(t, bulkResp)
   		
		search := apiClient.SearchAPI.Search(context.Background())
		searchRequest := openapiclient.NewSearchRequest("movies")
		testBasicSearch(search, searchRequest)			
	})

	fmt.Println("Search tests finished");
}
