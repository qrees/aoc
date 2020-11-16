from typing import List

class Solution:
    def _lexicalOrder(self, prefix):
        if prefix <= self.max:
            if prefix > 0:
                yield prefix
        else:
            return
        start = 1
        if prefix > 0:
            start = 0
        for i in range(start, 10):
            for res in self._lexicalOrder(prefix * 10 + i):
                yield res
    
    def lexicalOrder(self, n: int) -> List[int]:
        cur = 1
        self.max = n
        lista = []
        for res in self._lexicalOrder(0):
            lista.append(res)
        return lista


s = Solution()
print(s.lexicalOrder(156))
