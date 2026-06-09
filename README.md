# Pub/Sub

Projeto de testes para implementação de Pub/Sub usando o GCP.

- **Publisher** — produz mensagens sem saber quem vai consumir (`apps/api`)
- **Subscriber** — consome mensagens sem saber quem produziu (`apps/processor`)
- **Broker** — Google Cloud Pub/Sub faz o intermédio

O fluxo atual é:

```
apps/api → [appointments.created] → apps/processor → Mailtrap
```

Como cada subscription recebe uma cópia das mensagens, o mesmo tópico pode alimentar vários consumidores. Exemplo de
evolução possível:

```
                       ┌─► subscription: send-email   → processor de email
[appointments.created]─┤
                       └─► subscription: audit-log    → processor de auditoria
```

### Mensagem

Cada appointment criado pela API publica este payload no tópico:

```json
{
  "appointment_id": 1,
  "contact_name": "Matheus",
  "contact_email": "matheus@test.com",
  "scheduled_at": "2026-05-26T23:00:00+00:00"
}
```

O `processor` deserializa e usa esses dados para enviar o email de confirmação.
