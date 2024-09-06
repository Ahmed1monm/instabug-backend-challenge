class ChatsController < ApplicationController

  def index
    application = Application.find_by(token: params[:application_token])
    @chats = application.chats.find_by(application_id: application.id) if application
    if @chats
      render json: @chats, status: :ok
    else
      render json: { error: "application not found" }, status: :not_found unless application
      render json: [], status: :ok unless @chats
    end
  end

  def show
    application = Application.find_by(token: params[:application_token])
    @chat = application.chats.find_by(number: params[:number], application_id: application.id) if application
    if @chat
      render json: @chat, status: :ok
    else
      render json: { error: "Chat not found" }, status: :not_found
    end
  end

  def create
    application = Application.find_by(token: params[:application_token])
    @chat = application.chats.new(chat_params) if application

    if @chat && @chat.save
      render json: @chat, status: :created
    else
      render json: @chat.errors, status: :unprocessable_entity
    end
  end

  def update
    application = Application.find_by(token: params[:application_token])
    @chat = application.chats.find_by(number: params[:number])
    if @chat && @chat.update(chat_params)
      render json: @chat, status: :ok
    else
      render json: @chat.errors, status: :unprocessable
    end
  end

  private

  def chat_params
    params.require(:chat).permit(:name)
  end
end
