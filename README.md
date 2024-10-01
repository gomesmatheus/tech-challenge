## Criação do cluster na AWS
- `terraform init`
- `terraform plan`
- `terraform apply`

## Vincular kubectl com aws
- `aws eks --region $(terraform output -raw region) update-kubeconfig \ --name $(terraform output -raw cluster_name)`


## Executando via k8s
### Criação da base
#### PVS:
- `kubectl apply -f postgres-pv.yaml`
- `kubectl apply -f postgres-pvc.yaml`

#### Pods:
- `kubectl apply -f postgres-dep.yaml`
- `kubectl apply -f redis-dep.yaml`

#### Service:
- `kubectl apply -f postgres-service.yaml`
- `kubectl apply -f redis-service.yaml`

### Criação da APP
- `kubectl apply -f pod-go-app.yaml`
- `kubectl apply -f svc-go-app.yaml`

### Esse passo apenas deve ser executado caso necessário rodar os pods localmente (kubectl minikube)
Tive um problema com o minikube + wsl2 com a rede, portanto para acessar os pods de fora do cluster, tive que usar o port forward do kubectl, da seguinte forma:
- `kubectl port-forward go-app-6f7664cf8c-dlnfk 3333:3333` (Substituir **go-app-6f7664cf8c-dlnfk** por algum outro pod criado pelo deployment, checar com `kubectl get pods`)
- `kubectl port-forward postgres-dep-565c965f89-qkx99 5432:5432` (O mesmo vale para o **postgres-dep-565c965f89-qkx99**)

### Testando
Primeiro deve ser criado um cliente 
POST /cliente
```
{
    "cpf": 999999999,
    "nome": "João da Silva",
    "email": "joao@gma1il.com"
}
```

Depois os produtos
POST /produto
```
{
    "categoria_id": 2,
    "nome": "Coca Zero",
    "descricao": "Refrigerante",
    "preco": 5.75,
    "tempo_de_preparo": 1
}
```
```
{
    "categoria_id": 1,
    "nome": "X-burguer",
    "descricao": "Hmmm que delicia",
    "preco": 30.75,
    "tempo_de_preparo": 40
}
```

Depois podem ser criados pedidos variados com base nos produtos e nos clientes cadastrados
POST /pedido
```
{
    "cpf": 999999999,
    "metodo_de_pagamento": "Cartão",
    "produtos": [
        {
            "produto_id": 1,
            "quantidade": 1,
            "observacao": "sem molho"
        },
        {
            "produto_id": 2,
            "quantidade": 2,
            "observacao": "Sem acucar"
        }
    ]
}
```

Os pedidos podém ser listados
GET /pedido
Resposta:
```
[
    {
        "id": 1,
        "cpf": 999999999,
        "produtos": [
            {
                "produto_id": 1,
                "quantidade": 1,
                "observacao": "sem molho"
            },
            {
                "produto_id": 2,
                "quantidade": 2,
                "observacao": "Sem acucar"
            }
        ],
        "status": "Em preparo",
        "metodo_de_pagamento": "Cartão",
        "pagamento_aprovado": true
    }
]
```

Sendo possível alterar o status do pedido
PATCH /pedido/atualizar/{idPedido}
```
{
    "status": "Em preparo"
}
```

A app também possui um endpoint de webhook que simula a aprovação do pagamento
POST /pagamento/webhooks/
```
{
    "id": {{idPedidoConfirmado}}
}
```
