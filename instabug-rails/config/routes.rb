Rails.application.routes.draw do
  # Applications routes
  resources :applications, param: :token, only: [ :create, :show, :update, :index ] do
    # Nested routes for Chat within an Application
    resources :chats, param: :number, only: [ :create, :show, :index, :update ] do
      # Nested routes for Messages within a Chat
      resources :messages, param: :message_number, only: [ :create, :show, :index, :update ] do
        collection do
          get "search", to: "messages#search"
        end
      end
    end
  end

end
