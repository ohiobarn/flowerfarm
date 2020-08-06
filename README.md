# flowerfarm

The Ohio Barn Flower Farm, TOC:

[Docs](#Docs)

[Wordpress](#Wordpress)

---

## Docs

This repo uses [mkdocs](https://www.mkdocs.org/) ([help](https://mkdocs.readthedocs.io/en/0.10/)) and [github pages](https://help.github.com/articles/configuring-a-publishing-source-for-github-pages/) to host content at:

https://tonygilkerson.github.io/flowerfarm/

### Develop

```bash
#  one time setup for material theme
python3 -m pip install mkdocs-material

# then
mkdocs serve
# Edit content and review changes here:
open http://127.0.0.1:8000/
```

### Publish

```bash
cd OpsDoc
mkdocs build --clean
mkdocs gh-deploy
open https://tonygilkerson.github.io/flowerfarm/
```

