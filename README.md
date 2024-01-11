# MoneyFlowX

## Introduction

MoneyFlowX is a financial institution API that facilitates electronic wallet services for partners. This project, released under the MIT license, employs a diverse tech stack, including Gin, validator/10, jwt/j4, google/uuid, godotenv, crypto, viper, cast, gomail.v2, lumberjack.v2, gORM, and PostgreSQL.

## Key Aspects

- **Security First**: Robust security measures using HMAC-SHA1 for request body hashing and authentication through `X-UserId` and `X-Digest` headers.

- **Account Types**: Two distinct electronic wallet account types - identified and unidentified.

- **Balance Limits**: Clear maximum balance limits: 10,000.00 somoni for unidentified accounts and 100,000.00 somoni for identified accounts.

## Setup Guidelines

1. **Clone Repository**: Get the MoneyFlowX repository.
2. **Dependencies Installation**: Install dependencies specified in the tech stack.
3. **Environment Setup**: Configure environment variables seamlessly using `godotenv`.
4. **Database Configuration**: Easily set up the database by providing connection details in the `config.yaml` file.
5. **Run Your Application**: Start the application effortlessly for testing purposes.

## Contribution Opportunities

Contribute to the ongoing development of MoneyFlowX and share your valuable feedback. Ensure compliance with the MIT license terms.