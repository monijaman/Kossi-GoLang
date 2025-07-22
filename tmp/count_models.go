package main
import (
    "kossti/internal/infrastructure/database/models"
    "fmt"
)
func main() {
    models := []interface{}{
        &models.UserModel{}, &models.PasswordResetTokenModel{}, &models.SessionModel{}, &models.CacheModel{}, &models.CacheLockModel{},
        &models.JobModel{}, &models.JobBatchModel{}, &models.FailedJobModel{}, &models.PersonalAccessTokenModel{}, &models.HistoryLogModel{},
        &models.PermissionModel{}, &models.RoleModel{}, &models.ModelHasPermissionModel{}, &models.ModelHasRoleModel{}, &models.RoleHasPermissionModel{},
        &models.ProductModel{}, &models.ProductReviewModel{}, &models.CategoryModel{}, &models.BrandModel{}, &models.BrandCategoryModel{}, &models.ProductableModel{},
        &models.ProductTranslationModel{}, &models.CategoryTranslationModel{}, &models.BrandTranslationModel{}, &models.ProductReviewTranslationModel{}, &models.FeedbackTranslationModel{},
        &models.CommentModel{}, &models.CommentTranslationModel{},
        &models.SpecificationKeyModel{}, &models.SpecificationModel{}, &models.SpecificationTranslationModel{}, &models.SpecificationKeyTranslationModel{},
        &models.ImageModel{}, &models.TagModel{}, &models.FeedbackModel{}, &models.FormGeneratorModel{},
    }
    fmt.Printf("Total models migrated: %d\n", len(models))
}
