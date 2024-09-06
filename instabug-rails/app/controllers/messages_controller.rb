class MessagesController < ApplicationController

  def index
    application = Application.find_by(token: params[:application_token])
    unless application
      render json: { error: "application not found" }, status: :not_found
      return
    end

    chat = Chat.find_by(number: params[:chat_number], application_id: application.id)
    unless chat
      render json: { error: "chat not found" }, status: :not_found
      return
    end

    @messages = Message.where(chat_id: chat.id)
    if @messages.any?
      render json: @messages, status: :ok
    else
      render json: [], status: :ok
    end
  end


  def show
      application = Application.find_by(token: params[:application_token])
      unless application
        render json: { error: "application not found" }, status: :not_found
        return
      end

      chat = application.chats.find_by(number: params[:chat_number]) if application
      unless chat
        render json: { error: "chat not found" }, status: :not_found
        return
      end

      cache_key = "message/#{application.token}/#{chat.number}/#{params[:message_number]}"
      @message = $redis.get(cache_key)
      unless @message
        puts "Cache miss"
        @message = chat.messages.find_by(number: params[:message_number]) if chat
        if @message
          @message = @message.slice("body", "number", "created_at", "updated_at")
          $redis.set(cache_key, @message.to_json)
        end
      end

      if @message
        render json: @message, status: :ok
      else
        render json: { error: "Message not found" }, status: :not_found
      end
  end

    def create
      application = Application.find_by(token: params[:application_token])
      unless application
        render json: { error: "application not found" }, status: :not_found
        return
      end
      chat = application.chats.find_by(number: params[:chat_number]) if application
      unless chat
        render json: { error: "chat not found" }, status: :not_found
        return
      end

      @message = chat.messages.new(message_params) if chat
      if @message && @message.save
        render json: @message, status: :created
      else
        render json: @message.errors, status: :unprocessable_entity
      end
    end

    def update
      application = Application.find_by(token: params[:application_token])
      chat = application.chats.find_by(number: params[:chat_number]) if application
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
