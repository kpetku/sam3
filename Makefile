
USER_GH=eyedeekay
VERSION=0.0.2

version:
	gothub delete -s $(GITHUB_TOKEN) -u $(USER_GH) -r sam3 -t v$(VERSION) 2> /dev/null; true
	gothub release -s $(GITHUB_TOKEN) -u $(USER_GH) -r sam3 -t v$(VERSION) -d "version $(VERSION)"

