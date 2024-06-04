# Help
```
slidesgen --help 
slidesgen tpl --help
slidesgen gen --help
```


# slidesgen server(go version)
```
slidesgen -port 5001 -dir /path/to/images

visit
http://127.0.0.1:5001/index.html
```

# Basic
```
slidesgen tpl # generate a template
slidesgen gen # generate a single html file for the slides

```

# Advanced
```
slidesgen tpl -o myslides.md # generate a template with a custom name
slidesgen gen -i myslides.md # generate a single html file for the slides with a custom name
slidesgen gen -i myslides.md -o myslides.html # generate a single html file for the slides with a custom name and a custom output name
slidesgen gen -i myslides.md -o myslides.html -g 'data-background-gradient="linear-gradient(to bottom, #17b2c3,#283b95 )"' # generate a single html file for the slides with a custom name and a custom output name and a custom global options
```


# Example (slides.md)
```

```
