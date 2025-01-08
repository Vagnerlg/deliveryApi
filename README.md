# Estudos

A Ideia inicial é pensar em um fluxo de como um seguência de compra como um `pipeline`. Ex. uma compra de _pizza_ tem a escolha de _sabore(s)_ e tb opcionais de _borda_ alem de produtos adicionais como _bebidas_ e _sobremesas_, tudo isso pode ser um pipe de tipos de item.

> Questão de calculo desse `pipeline` pode ser complicado por ter item que se somem ou mesmo item que não tem valor.

> Esta logica terá muitos possibilidade e campos nulos que teram que ser resolvidos no front.

## Combo do mc


## Pizza
Um pedido de pizza com opção de escolha de até 3 sabores mais escolha de borda e adição de bebida e outros produtos.
```json
{
    "id": "676f19546e9f9c1f0b52d45a",
    "name": "Pizza",
    "description": "Com varios sabores",
    "category": "food",
    "tags": [
        "lanche",
    ],
    "pictures": [
        "uri1",
        "uri2",
    ],
    "price": 5490,
    "options": [
        {
            "id": "676f19546e9f9c1f0b52d45a",
            "kind": "options",
            "rules": {
                "min": 1,
                "max": 3,
                "calculated": "byQuota|byUnit"
            },
            "items": [
                {
                    "name": "Queijo",
                    "description": "muito queijo",
                    "pictures": [
                        "uri1",
                        "uri2",
                    ],
                    "price": 5490
                },
                {
                    "name": "Calabresa",
                    "description": "com rodelas de cebola",
                    "pictures": [
                        "uri1",
                        "uri2",
                    ],
                    "price": 5490
                }
            ] 
        },
        {
            "id": "dadad",
            "kind": "options",
            "rules": {
                "min": 1,
                "max": 1,
            },
            "items": [
                {
                    "name": "Simple",
                    "price": 5490
                },
                {
                    "name": "Com chedar",
                    "price": 5490
                }
            ]
        }
    ],
    "extra": [
        {
            "id": "dadad",
            "kind": "category",
            "categoryId": "sddsaaacas"
        }
    ]
}
```

## Esfiha
## Combo de japones
## Pratos comerciais
## Comida self-service