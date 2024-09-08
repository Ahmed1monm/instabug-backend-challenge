class ChatsController < ApplicationController

  def index
    application = Application.find_by(token: params[:application_token])
    
    if application.nil?
      render json: { error: "Application not found" }, status: :not_found
    else
      @chats = application.chats
      render json: @chats, status: :ok
    end
  end

  def show
    application = Application.find_by(token: params[:application_token])
    @chat = application.chats.find_by(number: params[:number], application_id: application.id) if application
    if @chat
      render json: @chat, status: :ok
    else
      render json: { error: "chat not found" }, status: :not_found
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
    
    if application.nil?
      render json: { error: "Application not found" }, status: :not_found
    else
      @chat = application.chats.find_by(number: params[:number])
      
      if @chat.nil?
        render json: { error: "Chat not found" }, status: :not_found
      elsif @chat.update(chat_params)
        render json: @chat, status: :ok
      else
        render json: @chat.errors, status: :unprocessable_entity
      end
    end
  end

  private

  def chat_params
    params.require(:chat).permit(:name)
  end
end
