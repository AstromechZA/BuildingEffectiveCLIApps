.PHONY: build 
build: slides.html

.PHONY: serve
serve: md-slides
	./md-slides serve -hot slides.md

slides.html: md-slides slides.md
	./md-slides serve --export-to=slides.html slides.md 

md-slides:
	$(shell curl https://raw.githubusercontent.com/AstromechZA/md-slides/master/install.sh -O)
	bash install.sh 
	rm install.sh
