Memoir BOT
_______

Memoir is a free software project which aims to help to manage all your life objectives.


INTEGRATIONS
____________


.. table:: chat tools we support
   :widths: auto

   ========= ============ =============== ============
     Slack     Telegram      RocketChat     Discord
   ========= ============ =============== ============
     TODO       DOING          TODO          TODO
   ========= ============ =============== ============


INSTALL
_______

For now, the only way to run Memoir is building from source. You can clone this git repo and the `memoir-gateway <https://github.com/luanguimaraesla/memoir-gateway>`_ project and run a few commands to get this prototype running.

1. Use go get to download the repositories
      .. code:: bash

         go get github.com/luanguimaraesla/memoir
         go get github.com/luanguimaraesla/memoir-gateway

2. Compile protobuf files
      .. code:: bash

         cd $GOPATH/src/github.com/luanguimaraesla/memoir
         protoc -I protobuff/ protobuff/metrics_gateway.proto --go_out=plugins=grpc:metricsgateway
         cd $GOPATH/src/github.com/luanguimaraesla/memoir-gateway
         protoc -I protobuff/ protobuff/metrics_gateway.proto --go_out=plugins=grpc:metricsgateway

3. Ensure dependencies, compile and install the programs
      .. code:: bash

         cd $GOPATH/src/github.com/luanguimaraesla/memoir
         dep ensure
         go install
         cd $GOPATH/src/github.com/luanguimaraesla/memoir-gateway
         dep ensure
         go install


USAGE
_____

Create your own manifest with questions you want to know about you. We only support numerical response values, so you need to think about how to create quantitative questions. Look at this example:

.. code:: yaml

  questions:
  - text: "Give a score to your happiness today."
    group: "happiness"
    metric: "daily_score"
    kind: "gauge"
    repeat: "daily"
    options: ["1", "2", "3", "4", "5", "6", "7", "8", "9"]

  - text: "How many tasks you finished at the job today?"
    group: "work"
    metric: "tasks_completed"
    kind: "gauge"
    repeat: "weekdays"
    options: []

Each question is composed by the following attributes:

* **text**: the text bot will tell you.
* **group**: the metric group will be the metric prefix (e.g. group.metric).
* **metric**: the metric name.
* **kind**: `prometheus metric types <https://prometheus.io/docs/concepts/metric_types/>`_. Supports "gauge", "counter", "summary", "histogram".
* **repeat**: frequency the bot will tell the questions. Can be "weekdays", "weekends", "daily" or any `crontab expression <https://crontab.guru/>`_.
* **options**: answer options, must be a list of floats or an empty list if you want to reply custom values.

Save this configuration file as **questions.yaml**. Then run metrics gateway (memoir-gateway).

.. code:: bash

  memoir-gateway run --collector :50051 --prometheus :9090

Create your own bot on Telegram with BotFather and copy it's access key. From another terminal, run the memoir server.

.. code:: bash

   memoir run --agent telegram --token <TELEGRAM_BOT_ACCESS_KEY> --config questions.yaml

Test your bot is running. Ask him "answer" on the chat. He'll reply "42" if there are no errors.

Now, you can type "ask" and answer each question you described on the manifest. Check the generated metrics on **localhost:9090/metrics**. You may use Prometheus to scrape this endpoint and Grafana to create dashboards to your life metrics.

I hope this project help you to be a better person. :)

DEVELOPMENT
___________

[TODO]
