
USER_GH=eyedeekay
VERSION=0.32.1

echo:
	@echo "type make version to do release $(VERSION)"

version:
	gothub release -s $(GITHUB_TOKEN) -u $(USER_GH) -r sam3 -t v$(VERSION) -d "version $(VERSION)"

