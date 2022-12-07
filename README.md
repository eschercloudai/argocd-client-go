# ArgoCD Client Library Generator

Argo's API is "documented" as an OpenAPI schema.
There are no clients available, so we should be able to generate one.

However, it appears the schema itself is totally broken, so we also take some liberties here and also patch it to make it actually work.
Ideally we should push these fixes upstream!
