Rails.application.routes.draw do
  resources :applications, param: :token, only: [ :create, :show, :update, :index ] do
    resources :chats, param: :number, only: [ :create, :show, :index, :update ] do
      resources :messages, param: :message_number, only: [ :create, :show, :index, :update ] do
        collection do
          get "search", to: "messages#search"
        end
      end
    end
  end

end
