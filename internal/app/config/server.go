package config

type Server struct {
	Port                string   `env:"PORT"`
	Debug               bool     `env:"SERVER_DEBUG"`
	ReadTimeoutSeconds  int      `env:"SERVER_READ_TIMEOUT"`
	WriteTimeoutSeconds int      `env:"SERVER_WRITE_TIMEOUT"`
	UserTokenName       string   `env:"USER_TOKEN_NAME"`
	S3Bucket            S3Bucket `env:"AWS_S3_BUCKET"`
}

type S3Bucket struct {
	AWS_REGION        string `env:"AWS_REGION"`
	AWS_ACCESS_KEY_ID string `env:"AWS_ACCESS_KEY_ID"`
	AWS_SECRET_KEY    string `env:"AWS_SECRET_KEY"`
	BUCKET_NAME       string `env:"BUCKET_NAME"`
	S3_FOLDER_NAME    string `env:"S3_FOLDER_NAME"`
}
