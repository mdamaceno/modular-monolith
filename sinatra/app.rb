# frozen_string_literal: true

require_relative 'loader'

require 'sinatra/base'
require 'sinatra/contrib/all'
require 'json'

class Application < Sinatra::Base
  register Sinatra::Contrib

  include Components
  include Helpers

  before do
    content_type :json
  end

  get '/' do
    { message: 'The sound of silence...' }.to_json
  end

  namespace '/marketing' do
    get '/announces' do
      Serializer.serialize(Marketing.index)
    end

    post '/announces' do
      Serializer.serialize(Marketing.announce_something(request: request))
    end
  end

  post '/orders' do
    Serializer.serialize(Order.create_order(request: request))
  end
end
