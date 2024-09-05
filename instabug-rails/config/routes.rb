Rails.application.routes.draw do
  # Applications routes
  resources :applications, param: :token, only: [:create, :show, :update, :index] do
    # Nested routes for Chats within an Application
    resources :chats, param: :number, only: [:create, :show, :index] do
      # Nested routes for Messages within a Chat
      resources :messages, param: :message_number, only: [:create, :show] do
        collection do
          get 'search', to: 'messages#search'
        end
      end
    end
  end

end
