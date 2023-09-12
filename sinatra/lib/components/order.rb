module Components
  module Order
    def self.create_order(request: nil)
      puts request.inspect
      UseCase::CreateOrder.new.call
    end
  end
end
