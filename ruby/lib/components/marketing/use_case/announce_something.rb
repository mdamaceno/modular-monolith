module Components::Marketing::UseCase
  class AnnounceSomething
    def call
      string = 'Announcing something!'

      puts string

      string
    end
  end
end
