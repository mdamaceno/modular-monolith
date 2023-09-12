module Components
  module Marketing
    def self.index
      UseCase::AnnounceSomething.new.call
    end

    def self.announce_something(request: nil)
      puts request.inspect
      UseCase::AnnounceSomething.new.call
    end
  end
end
