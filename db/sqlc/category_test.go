package db

import (
	"context"
	"testing"

	"github.com/Juan-Ibanezdf/gofinance-backend/util"
	"github.com/stretchr/testify/require"
)

func createRandomCategory(t *testing.T) Category {
	user := createRandomUser(t)
	arg := CreateCategoryParams{
		UserID:      user.ID,
		Title:       util.RandomString(12),
		Type:        "debit",
		Description: util.RandomString(20),
	}

	category, err := testQueries.CreateCategory(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, category)

	require.Equal(t, arg.UserID, category.UserID)
	require.Equal(t, arg.Title, category.Title)
	require.Equal(t, arg.Type, category.Type)
	require.Equal(t, arg.Description, category.Description)
	require.NotEmpty(t, category.CreatedAt)

	return category
}

func TestCreateCategory(t *testing.T) {
	createRandomCategory(t)
}

func TestGetCategory(t *testing.T) {
	category1 := createRandomCategory(t)
	category2, err := testQueries.GetCategory(context.Background(), category1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, category2)

	require.Equal(t, category1.ID, category2.ID)
	require.Equal(t, category1.Title, category2.Title)
	require.Equal(t, category1.Description, category2.Description)
	require.Equal(t, category1.Type, category2.Type)
	require.NotEmpty(t, category2.CreatedAt)
}

func TestDeleteCategory(t *testing.T) {
	Category := createRandomCategory(t)
	err := testQueries.DeleteCategory(context.Background(), Category.ID)
	require.NoError(t, err)
}

func TestUpdateCategory(t *testing.T) {
	Category1 := createRandomCategory(t)

	arg := UpdateCategoryParams{
		ID:          Category1.ID,
		Title:       util.RandomString(12),
		Description: util.RandomString(20),
	}

	Category2, err := testQueries.UpdateCategory(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, Category2)

	require.Equal(t, Category1.ID, Category2.ID)
	require.Equal(t, arg.Title, Category2.Title)
	require.Equal(t, arg.Description, Category2.Description)
	require.NotEmpty(t, Category2.CreatedAt)
}

func TestListCategories(t *testing.T) {
	lastCategory := createRandomCategory(t)

	arg := GetCategoriesParams{
		UserID:      lastCategory.UserID,
		Type:        lastCategory.Type,
		Title:       lastCategory.Title,
		Description: lastCategory.Description,
	}

	categorys, err := testQueries.GetCategories(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, categorys)

	for _, category := range categorys {
		require.Equal(t, lastCategory.ID, category.ID)
		require.Equal(t, lastCategory.UserID, category.UserID)
		require.Equal(t, lastCategory.Title, category.Title)
		require.Equal(t, lastCategory.Description, category.Description)
		require.NotEmpty(t, lastCategory.CreatedAt)
	}
}
