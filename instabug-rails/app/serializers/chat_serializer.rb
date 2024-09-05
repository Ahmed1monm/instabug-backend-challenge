class ChatSerializer < ActiveModel::Serializer
  attributes :name, :number, :created_at, :updated_at
end
