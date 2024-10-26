package service

// Service is a identifier which service throw the error
// mainly used to determine whether internal application / external
type Service string

// General service which SkorLife app using
// if any other service we are using; with Service type we can extend it
// It's recommended to add in the package level

const (
	Application Service = "APPLICATION"
	Job         Service = "JOB"

	PostgresDB Service = "POSTGRES"
	Redis      Service = "REDIS"
	MongoDB    Service = "MONGO"

	JIRA Service = "JIRA"

	MailCow Service = "MAILCOW"
	GOOGLE  Service = "GOOGLE"

	VIDA     Service = "VIDA"
	AKSATA   Service = "AKSATA"
	INFO_BIP Service = "INFO_BIP"

	CLIK Service = "CLIK"
	CBI  Service = "CBI"

	AWSCloud       Service = "AWS"
	AWSCloudLambda Service = "AWS_LAMBDA"
	AWSCloudS3     Service = "AWS_S3"
	AWSCloudSES    Service = "AWS_SES"
	AWSCloudKMS    Service = "AWS_KMS"
)
