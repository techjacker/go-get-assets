test:
	@go build && go-get-assets -o $(PWD)/cms.output.json -r /images $(PWD)/fixtures/cms.json $(PWD)/images

.PHONY: test