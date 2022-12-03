###############################################################################
#                                                                             #
# Адаптер - это структурный паттерн, который позволяет подружить              #
# несовместимые объекты.                                                      #
#                                                                             #
# Адаптер выступает прослойкой между двумя объектами, превращая вызовы одного # 
# в вызовы понятные другому.                                                  #
#                                                                             #
###############################################################################


class Target:
    """
    Целевой класс объявляет интерфейс, с которым может работать клиентский код.
    """
    
    def request(self) -> str:
        return "Target: The default target's behavior."


class Adaptive:
    """
    Адаптируемый класс содержит некоторое полезное поведение, но его интерфейс 
    несовместим с существующим клиентским кодом. Адаптируемый класс нуждается в 
    некоторой доработке, прежде чем клиентский код сможет его использовать.
    """

    def specific_request(self) -> str:
        return ".evitpadA eht fo roivaheb laicepS"


class Adapter(Target):
    """
    Адаптер делает интерфейс Адаптируемого класса совместимым с целевым 
    интерфейсом благодаря агрегации.
    """

    def __init__(self, adaptive: Adaptive) -> None:
        self.adaptive = adaptive

    def request(self) -> str:
        return f"Adapter: (TRANSLATED) {self.adaptive.specific_request()[::-1]}"


def client_code(target: Target) -> None:
    """
    Клиентский код поддерживает все классы, использующие интерфейс Target.
    """

    print(target.request(), end="")


if __name__ == "__main__":
    print("Client: I can work just fine with the Target objects:")
    target = Target()
    client_code(target)
    print("\n")

    adaptive = Adaptive()
    print("Client: The Adaptive class has a weird interface."
          "See, I don't understand it:")
    print(f"Adaptive: {adaptive.specific_request()}", end="\n\n")

    print("Client: But I can work with it via the Adapter:")
    adapter = Adapter(adaptive)
    client_code(adapter)

    # Client: I can work just fine with the Target objects:
    # Target: The default target's behavior.

    # Client: The Adaptive class has a weird interface.See, I don't understand it:
    # Adaptive: .evitpadA eht fo roivaheb laicepS

    # Client: But I can work with it via the Adapter:
    # Adapter: (TRANSLATED) Special behavior of the Adaptive.
