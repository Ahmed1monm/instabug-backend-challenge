class ChatController < ApplicationController
  def index
    @chats = Chat.all
    render json: @chats, status: :ok
  end

  def show
    @chat = Chat.find(params[:id])
    render json: @chat, status: :ok
  end

  def create
    @chat = Chat.new(chat_params)

    if @chat.save
      render json: @chat, status: :created
    else
      render json: @chat.errors, status: :unprocessable_entity
    end
  end

  private

  def chat_params
    params.require(:chat).permit(:name)
  end
end
