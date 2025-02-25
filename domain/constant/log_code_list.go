package constant

const (
	UnlockPDFByStatementIDDBFailureErr = "ERR.STATEMENT.SERVICES.UNLOCK_PDF_BY_STATEMENT_ID.DB_FAILURE"
	UnlockPDFByStatementIDDBCommitErr  = "ERR.STATEMENT.SERVICES.UNLOCK_PDF_BY_STATEMENT_ID.DB_COMMIT"

	UnlockPDFByStatementIDSuccessInfo = "INFO.STATEMENT.SERVICES.UNLOCK_PDF_BY_STATEMENT_ID.SUCCESS"

	UploadPDFBatchDBFailureErr = "ERR.STATEMENT.SERVICES.UPLOAD_PDF_BATCH.DB_FAILURE"
	UploadPDFBatchUploadErr    = "ERR.STATEMENT.SERVICES.UPLOAD_PDF_BATCH.UPLOAD"
	UploadPDFBatchFailureErr   = "ERR.STATEMENT.SERVICES.UPLOAD_PDF_BATCH.FAILURE"
	UploadPDFBatchUpdateErr    = "ERR.STATEMENT.SERVICES.UPLOAD_PDF_BATCH.UPDATE"
	UploadPDFBatchDBCommitErr  = "ERR.STATEMENT.SERVICES.UPLOAD_PDF_BATCH.DB_COMMIT"

	CheckExtractedPDFByPDFIDTransactionPatternsDBFailureErr = "ERR.STATEMENT.SERVICES.CHECK_EXTRACTED_PDF_BY_PDF_ID.TRANSACTION_PATTERNS.DB_FAILURE"
	CheckExtractedPDFByPDFIDDBFailureErr                    = "ERR.STATEMENT.SERVICES.CHECK_EXTRACTED_PDF_BY_PDF_ID.DB_FAILURE"

	CheckExtractedPDFByIDTransactionPattersAmountInfo = "INFO.STATEMENT.SERVICES.CHECK_EXTRACTED_PDF_BY_PDF_ID.TRANSACTION_PATTERS_AMOUNT"
	CheckExtractedPDFByIDSuccessInfo                  = "INFO.STATEMENT.SERVICES.CHECK_EXTRACTED_PDF_BY_PDF_ID.SUCCESS"

	CheckExtractedPDFBatchSuccessAmountInfo = "INFO.STATEMENT.SERVICES.CHECK_EXTRACTED_PDF_BATCH.SUCCESS_AMOUNT"

	GetOneAlreadyUnlockedPDFByStatementIDDBFailureErr = "ERR.STATEMENT.SERVICES.GET_ONE_ALREADY_UNLOCKED_PDF_BY_STATEMENT_ID.DB_FAILURE"
	GetOneAlreadyUnlockedPDFByStatementIDDBNotFoundInfo = "INFO.STATEMENT.SERVICES.GET_ONE_ALREADY_UNLOCKED_PDF_BY_STATEMENT_ID.DB_NOT_FOUND"
	GetOneAlreadyUnlockedPDFByStatementIDSuccessInfo = "INFO.STATEMENT.SERVICES.GET_ONE_ALREADY_UNLOCKED_PDF_BY_STATEMENT_ID.SUCCESS"

	LibHTTPRequestInfo  = "INFO.STATEMENT.LIB.HTTP.REQUEST"
	LibHTTPResponseInfo = "INFO.STATEMENT.LIB.HTTP.RESPONSE"
)
