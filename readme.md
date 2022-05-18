#Products-API

## Live Environment
This API is hosted using netlify free tier and allows only 10 requests per day. This is done to save my monthly quota on Netlify. In case this is not sufficient, please write to ![me](https://gfycat.com/pointedvacanteskimodog)

## Pre-Requesite

API can be run locally on any machine that satisfy below pre-requesites
- Go 1.18
- make software

## Local setup

Follow below steps to locally setup Products-API
- clone this repository
- run `make run`
- visit http://localhost:1234/swagger/index.html

## Running tests

To run test run, `make test` . This would run all unit tests 

## Rate limits

API runs with requests rate limits. Currently the values are set to allow 10 requests per day and follows token bucket rate limit alogorithm.
