defmodule WsHandler do
  @behaviour :cowboy_websocket

  def init(req, state) do
    {:cowboy_websocket, req, state}
  end

  def websocket_init(state) do
    {:ok, state}
  end

  def websocket_handle({:text, "hello"}, state) do
    {:reply, {:text, "world!"}, state}
  end

  def websocket_handle({:text, msg}, state) do
    {:reply, {:text, "echo: #{msg}"}, state}
  end

  def websocket_handle(_, state) do
    {:ok, state}
  end

  def websocket_info(_info, state) do
    {:ok, state}
  end

  def terminate(_reason, _req, _state) do
    :ok
  end
end
