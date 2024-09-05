class MessagesController < ApplicationController

    def index
      application = Application.find_by(token: params[:application_token])
      chat = Chat.find_by(number: params[:chat_number]) if application
      @messages = Message.all if chat

      if @messages
        render json: @messages, status: :ok
      else
        render json: { error: "Messages not found" }, status: :not_found
      end


    end

    def show
      # Find the application by token
      application = Application.find_by(token: params[:application_token])

      # Find the chat by number within the scope of the application
      chat = application.chats.find_by(number: params[:chat_number]) if application

      # Find the message by number within the scope of the chat
      @message = chat.messages.find_by(number: params[:message_number]) if chat

      if @message
        render json: @message, status: :ok
      else
        render json: { error: "Message not found" }, status: :not_found
      end
    end

    def create
      application = Application.find_by(token: params[:application_token])
      chat = Chat.find_by(number: params[:chat_number]) if application
      @message = chat.messages.new(message_params) if chat

      if @message && @message.save
        render json: @message, status: :created
      else
        render json: @message.errors, status: :unprocessable_entity
      end
    end

    def update
      application = Application.find_by(token: params[:application_token])
      chat = Chat.find_by(number: params[:chat_number]) if application
      @message = chat.messages.find_by(number: params[:message_number]) if chat
      if @message && @message.update(message_params)
        render json: @message, status: :ok
      else
        render json: @message.errors, status: :unprocessable
      end
    end

    private

    def message_params
      params.require(:message).permit(:body)
    end
end
