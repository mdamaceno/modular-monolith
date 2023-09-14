# frozen_string_literal: true

require_relative 'loader'

require 'sinatra/base'
require 'sinatra/contrib/all'
require 'json'

class Application < Sinatra::Base
  register Sinatra::Contrib

  before do
    content_type :json
  end

  namespace '/marketing' do
    extend Components::Marketing::Controller
  end

  namespace '/orders' do
    extend Components::Order::Controller
  end
end
