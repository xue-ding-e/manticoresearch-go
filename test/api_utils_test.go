/*
Manticore Search Client

Testing UtilsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"fmt"
	openapiclient "github.com/manticoresoftware/manticoresearch-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_openapi_UtilsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	configuration.Servers[0].URL = fmt.Sprintf("http://%s:9408", openapiclient.GetDefaultIP())
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test UtilsAPIService Sql", func(t *testing.T) {

		//t.Skip("skip test")  // remove to run test

		var sql string

		sql = "DROP TABLE IF EXISTS products"
		resp, httpRes, err := apiClient.UtilsAPI.Sql(context.Background()).Body(sql).RawResponse(true).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		fmt.Printf("%v\n", resp)

		assert.Equal(t, 200, httpRes.StatusCode)

		sql = "CREATE TABLE IF NOT EXISTS products (title text, price float, sizes multi, meta json, coeff float, tags1 multi, tags2 multi)"
		resp, httpRes, err = apiClient.UtilsAPI.Sql(context.Background()).Body(sql).RawResponse(true).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		fmt.Printf("%v\n", resp)

		assert.Equal(t, 200, httpRes.StatusCode)

		sql = "SELECT * FROM products"
		resp, httpRes, err = apiClient.UtilsAPI.Sql(context.Background()).Body(sql).RawResponse(false).Execute()
		require.Nil(t, err)
		require.NotNil(t, resp)
		fmt.Printf("%v\n", resp)

		assert.Equal(t, 200, httpRes.StatusCode)

		sql = "SELECT * FROM products"
		resp, httpRes, err = apiClient.UtilsAPI.Sql(context.Background()).Body(sql).RawResponse(true).Execute()
		require.Nil(t, err)
		require.NotNil(t, resp)
		fmt.Printf("%v\n", resp)

		assert.Equal(t, 200, httpRes.StatusCode)

		sql = "SELECT * FROM products"
		resp, httpRes, err = apiClient.UtilsAPI.Sql(context.Background()).Body(sql).Execute()
		require.Nil(t, err)
		require.NotNil(t, resp)
		fmt.Printf("%v\n", resp)

		assert.Equal(t, 200, httpRes.StatusCode)

		sql = "TRUNCATE TABLE products"
		resp, httpRes, err = apiClient.UtilsAPI.Sql(context.Background()).Body(sql).Execute()
		require.Nil(t, err)
		require.NotNil(t, resp)
		fmt.Printf("%v\n", resp)

		assert.Equal(t, 200, httpRes.StatusCode)

		sql = "SHOW TABLES"
		resp, httpRes, err = apiClient.UtilsAPI.Sql(context.Background()).Body(sql).Execute()
		require.Nil(t, err)
		require.NotNil(t, resp)
		fmt.Printf("%v\n", resp)

		assert.Equal(t, 200, httpRes.StatusCode)

		fmt.Println("Util tests finished")

	})

}
