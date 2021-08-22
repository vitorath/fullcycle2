Producer --> Broker(3) <- Consumer

Zookepper auxilia o kafka (por hora)

Topics

Canal de comunicacao responsavel por receber e e disponibilizar os dados enviados para o Kafka
Kafka nao trabalha igual ao rabbitmq, as filas do rabbbitmq eh mais ou menos que um topico no kafka


Producer --Topic--> Kafka --Topic--> Consumer 1
                          --Topic--> Consumer 1    

Topic eh mais ou menos como se fosse um log
A mensagem eh armazenada uma atras da outra e recebe um offset
Um consumer por estar sendo uma mensagem de um offset e outro de outro offset em um mesmo topic
E as mensagem pode ser relidas caso necessario pois sao armazenadas

Anatomia do registro
Offset 0
    Headers (Opt)
    Key   
    Value
    Timestamp  

Partitions
Cada ttopic pode ter uma ou mais particoes para conseguir garantir a distribuicao e resiliencia de seus dados

Topic --Message 1--> Partition 1 (Broker 1)
      --Message 2--> Partition 2 (Broker 22)
      --Message 3--> Partition 3 (Broker 1)

Efeito colaterial ao trabalhar com paricoes. Como garantir a ordem?
Partition 1 [0, 1] (Trasnferencia) <-- Consumer 1 (Slow)
Partition 2 [0] (Estorno) <-- Consumer 2 (Fast)
Partition 3 [0] <-- Consumer 2

Problema consumidor 1 est'a lento e o consumidor 2 esta rapido, logo tem chance do estorno vir antes da transferencia

Resolucao
Partition 1  (Trasnferencia[0] e estorno (1)) <-- Consumer 1 (Slow)
Partition 2 [0] <-- Consumer 2 (Fast)
Partition 3 [0] <-- Consumer 2

Para fazer com que as mensagens vao para mesma particao, para isso usamos as keys e com isso ira ser redicionado para mesma particao, caso as keys sejam iguais
Partition 1 [0, 1]
Trasnferencia (0) -> Key=Movimentacao
Estorno(1) > Key=Movimentacao

Particoes distribuidas

Topic Vendas -> Broker A (Partition 1)
             -> Broker B (Partition 2)
             -> Broker C (Partition 3)

Replicator Factor = 2
Topic Vendas -> Broker A (Partition 1, Partition 3)
             -> Broker B (Partition 2, Partition 1)
             -> Broker C (Partition 3, Partition 2)

Partition leadership
Consumidor faz a leitura na particao leader e o producer entrega ao leader
Parition Follow sao as elegiveis para ser leader
Topic Vendas -> Broker A (Partition 1[L], Partition 3)
             -> Broker B (Partition 1, Partition 2[L]) -- Caindo
             -> Broker C (Partition 3[L], Partition 2)

Topic Vendas -> Broker A (Partition 1[L], Partition 3)
             -> Broker B (Partition 1, Partition 2) -- Caiu
             -> Broker C (Partition 3[L], Partition 2[L])

Garantir entrega: 

O producer nao recebe a notificacao do leader de que a mensagem foi salva
Vantagem: processa mais mensagens
Desvantagem: corre risco de perder mensagens
Producer --Ack 0 [None]--> Broker A [Leader]
                           Broker B [Follower]
                           Broker C [Follower]


O leader salva e notifica de que salvou a mensagem ao producer
Vantagem: processa mais mensagens (menos que o cenario anterior)
Desvantagem: corre risco de perder mensagens, pois ainda nao replicou aos followers
Producer --Ack 1 [Leader]--> Broker A [Leader]
                             Broker B [Follower]
                             Broker C [Follower]

O Leader recebe envia para os followers, os followers notifica o leader de que foi salva e o leader notifica o producer
Vantagem: garante 100% de que a mensagem sera salva
Desvantagem: demora mais para processar as mensagens
Producer --Ack 1 [ALL]--> Broker A [Leader]
                          Broker B [Follower]
                          Broker C [Follower]

Garantir entrega (2): 
At most once: Melhor performance. Pode perder algumas mensagens
1,2,3,4,5 --> Kafka Process --> 1,3,4

At least once: Performance moderada.Prode duplicar mensagens
1,2,3,4,5 --> Kafka Process --> 1,2,3,4,4,5

Excly once: Pior perfomance. Exatamente um vez
1,2,3,4,5 --> Kafka Process --> 1,2,3,4,5

Producer Indepotencia
Cai a conexao e houve um retry e duplica uma mensagem
Indepotencia: OFF
Gravar duplicada
Consumer vai ler a mesma mensagem duas vezes

Indepotencia: ON
Kafka vai descartar uma das mensagens e garante que caia na ordem correta

Desvantagem: mais lentidao

Consumer
Caso nao tenha fgrupo o proprio consumidor eh considerado o grupo dele
Producer --> Topic [Partition 1] --> Consumer
         --> Topic [Partition 2] -->
         --> Topic [Partition 3] -->

Consumer groups
Caso defina um grupo, os consumos sao distribuidos entre os cosumers do grupo
Producer --> Topic [Partition 1] --> Consumer A [Group X]
         --> Topic [Partition 2] --> Consumer A [Group X]
         --> Topic [Partition 3] --> Consumer B [Group X]

Producer --> Topic [Partition 1] --> Consumer A [Group X]
         --> Topic [Partition 2] --> Consumer B [Group X]
         --> Topic [Partition 3] --> Consumer C [Group X]

Nao tem como dois consumidores no mesmo grupo ler a mesma particao, caso tenha consumers a mais que particoes, o consumer ficara parado
Producer --> Topic [Partition 1] --> Consumer A [Group X]
         --> Topic [Partition 2] --> Consumer B [Group X]
         --> Topic [Partition 3] --> Consumer C [Group X]
                                     Consumer D [Group X] (Idle)