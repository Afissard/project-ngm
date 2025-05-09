defmodule Server.Application do
  # See https://hexdocs.pm/elixir/Application.html
  # for more information on OTP Applications
  @moduledoc false

  use Application

  @impl true
  def start(_type, _args) do
    dispatch = :cowboy_router.compile([
      {:_, [
        {"/ws", WsHandler, []}
      ]}
    ])

    {:ok, _} = :cowboy.start_clear(:http,
      [{:port, 4000}],
      %{env: %{dispatch: dispatch}}
    )

    children = [
      # Starts a worker by calling: Server.Worker.start_link(arg)
      # {Server.Worker, arg}
    ]

    # See https://hexdocs.pm/elixir/Supervisor.html
    # for other strategies and supported options
    opts = [strategy: :one_for_one, name: Server.Supervisor]
    Supervisor.start_link(children, opts)
  end
end
