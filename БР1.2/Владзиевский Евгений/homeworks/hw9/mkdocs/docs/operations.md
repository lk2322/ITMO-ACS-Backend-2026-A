# Эксплуатация

## Быстрая диагностика

```bash
# состояние всех подов
kubectl -n rental get pods

# что произошло в кластере за последнее время
kubectl -n rental get events --sort-by=.lastTimestamp | tail -20

# логи сервиса
kubectl -n rental logs deployment/auth-service --tail=50 -f

# какая версия выкачена
kubectl -n rental get deployments \
  -o custom-columns='DEPLOY:.metadata.name,IMAGE:.spec.template.spec.containers[0].image'
```

## Типовые ситуации

### Под перезапускается по кругу

```bash
kubectl -n rental describe pod <имя>      # причина в Events
kubectl -n rental logs <имя> --previous   # логи упавшего контейнера
```

Чаще всего — недоступная база или неверная переменная окружения.

### Запросы возвращают 401, хотя токен свежий

Проверьте, что сервису передан `JWT_SECRET`:

```bash
kubectl -n rental get deployment property-service \
  -o jsonpath='{.spec.template.spec.containers[0].env[*].name}'
```

Сервисы `property-service`, `payment-service`, `messaging-service`
проверяют подпись сами и без общего ключа отвергают любые токены.

### Нужно откатить версию

```bash
kubectl -n rental rollout undo deployment/api-gateway
kubectl -n rental rollout status deployment/api-gateway
```

Или выкатить конкретный тег через GitHub Actions:
**Run workflow** → поле `image_tag` → короткий SHA нужного коммита.

### Проверить события в Kafka

```bash
kubectl -n rental exec kafka-0 -- kafka-topics \
  --bootstrap-server kafka:9092 --list

kubectl -n rental exec kafka-0 -- kafka-console-consumer \
  --bootstrap-server kafka:9092 --topic user.events \
  --from-beginning --max-messages 5
```

### Посмотреть данные сервиса

```bash
kubectl -n rental exec auth-db-0 -- \
  psql -U auth -d auth_db -c "select id, email, role from users limit 10;"

# не опубликованные события outbox
kubectl -n rental exec auth-db-0 -- \
  psql -U auth -d auth_db -c "select * from outbox_events where status <> 'published';"
```

## Регулярные задачи

| Что | Как часто | Команда |
|---|---|---|
| Проверить свободное место | еженедельно | `df -h /` |
| Почистить старые образы | ежемесячно | `sudo docker image prune -a` |
| Проверить доступность сервисов | постоянно | дашборд Grafana |
