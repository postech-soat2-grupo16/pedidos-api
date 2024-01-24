Feature: API Pedido

  Scenario Outline: Pedido creation
    Given Parameter ClientID: <clientID>
    When request POST /pedido
    Then statusCode should be <statusCode>

    Examples:
        | clientID | statusCode |
        | 21       | 201       |

  Scenario Outline: Pedido update
    Given get first order ID
    When request PUT /pedido with status "<status>"
    Then statusCode should be <statusCode>

    Examples:
      | statusCode | status |
      | 200       | CRIADO |
      | 422       | ERRO   |

  Scenario Outline: Pedido patch
    Given get first order ID
    When request PATCH /pedido with status "<status>"
    Then statusCode should be <statusCode>

    Examples:
      | statusCode | status |
      | 200       | CRIADO |
      | 422       | ERRO   |

  Scenario Outline: Pedido GET - Success
    Given get first order ID
    When request GET /pedido by id
    Then statusCode should be <statusCode>

    Examples:
      | statusCode |
      |  200       |


  Scenario Outline: Pedido GET - Error
    Given unknown order ID
    When request GET /pedido by id
    Then statusCode should be <statusCode>

    Examples:
      | statusCode |
      |  404       |

  Scenario Outline: Pedido GET - With clientID
    Given Parameter ClientID: <clientID>
    When request GET /pedido with ClientID
    Then statusCode should be <statusCode>

    Examples:
      | clientID | statusCode |
      | 21       | 200       |
      | 404       | 200       |

  Scenario Outline: Pedido GET - With status
    Given Parameter Status: "<status>"
    When request GET /pedido with Status
    Then statusCode should be <statusCode>

    Examples:
      | status | statusCode |
      | CRIADO | 200       |
      | ERRO | 200       |

  Scenario Outline: Should get healthcheck
    When request GET /healthcheck
    Then statusCode should be <statusCode>

    Examples:
      | statusCode |
      | 200       |

  Scenario Outline: Should delete order
    Given get first order ID
    When request DELETE /pedido by id
    Then statusCode should be <statusCode>

    Examples:
      | statusCode |
      |  204       |
