# Сборка статического сайта документации и его раздача nginx.
# Контекст сборки — каталог hw9, чтобы в образ попала и спецификация OpenAPI.

FROM squidfunk/mkdocs-material:9.5.44 AS build
WORKDIR /docs
COPY mkdocs/mkdocs.yml ./mkdocs.yml
COPY mkdocs/docs ./docs
# спецификация лежит рядом с документацией: Swagger UI тянет её с того же сайта
COPY openapi/openapi.yaml ./docs/openapi.yaml
RUN mkdocs build --strict --site-dir /site

FROM nginx:alpine
# Ingress отдаёт путь вместе с префиксом /docs, поэтому сайт кладётся
# в соответствующий подкаталог — иначе nginx ищет /docs/index.html и не находит
COPY --from=build /site /usr/share/nginx/html/docs
EXPOSE 80
