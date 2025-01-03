# Monitoramento de Consumo de Água - API

## **Descrição**
Esta API permite monitorar o consumo diário de água de acordo com as necessidades individuais do usuário. Com base em dados como peso e nível de atividade física, a API calcula a quantidade ideal de água que o usuário deve consumir diariamente. O sistema utiliza autenticação JWT para segurança.

---

## **Recursos da API**

### **Autenticação**
- Registro de usuários.
- Login e logout com autenticação JWT.
- Tokens de acesso e refresh para segurança.

### **Gerenciamento de Perfil**
- Inserção e atualização de informações pessoais:
  - Peso corporal.
  - Nível de atividade física (baixo, moderado, alto).
  - Meta personalizada de consumo diário de água (opcional).

### **Cálculo da Necessidade Diária de Água**
- Fórmula base: `35 mL por kg de peso corporal`.
- Ajustes automáticos pelo nível de atividade física:
  - Baixo: +0%.
  - Moderado: +10%.
  - Alto: +20%.

### **Registro de Consumo**
- Registro da quantidade de água consumida em mL.
- Rastreamento do progresso diário em relação à meta.

### **Relatórios e Histórico**
- Exibição do resumo diário.
- Acesso ao histórico semanal e mensal.

---

## **Regras de Negócio**
1. O consumo diário máximo permitido é de 4 litros.
2. Usuários podem configurar uma meta personalizada que substituirá o cálculo automático.
3. Tokens JWT expiram em 24 horas e podem ser renovados com um token de refresh.
4. Um registro de consumo deve conter:
   - Data e hora.
   - Quantidade consumida (em mL).
5. Registros podem ser alterados ou removidos apenas no mesmo dia.

---

## **Endpoints Principais**

### **Autenticação**
- **POST** `/register` - Registrar novo usuário.
- **POST** `/login` - Realizar login e obter token JWT.
- **POST** `/refresh` - Renovar token de acesso.
- **POST** `/logout` - Invalidar sessão.

### **Perfil do Usuário**
- **GET** `/profile` - Obter dados do perfil.
- **PUT** `/profile` - Atualizar informações pessoais.

### **Consumo de Água**
- **GET** `/water-intake` - Obter progresso diário e meta.
- **POST** `/water-intake` - Registrar consumo de água.
- **PUT** `/water-intake/{id}` - Atualizar registro de consumo.
- **DELETE** `/water-intake/{id}` - Excluir registro de consumo.

### **Histórico**
- **GET** `/history` - Obter histórico semanal ou mensal.

---

## **Instalação e Configuração**
1. Clone o repositório:
   ```bash
   git clone https://github.com/eduardonakaidev/drink-water-api.git
   ```
2. Navegue até o diretório do projeto:
   ```bash
   cd drink-water-api
   ```
3. Instale as dependências:
   ```bash
   npm install
   ```
4. Configure as variáveis de ambiente no arquivo `.env`:
   ```env
   PORT=3000
   JWT_SECRET=sua_chave_secreta
   DB_URI=seu_banco_de_dados
   ```
5. Inicie o servidor:
   ```bash
   npm start
   ```

---

## **Tecnologias Utilizadas**
- Golang
- http/net
- gorm
- JSON Web Token (JWT)
- Postgres

---

## **Contribuição**
Contribuições são bem-vindas! Siga os passos abaixo:
1. Faça um fork do repositório.
2. Crie uma branch para sua feature/fix:
   ```bash
   git checkout -b minha-feature
   ```
3. Faça os commits:
   ```bash
   git commit -m "Descrição do commit"
   ```
4. Envie para o repositório:
   ```bash
   git push origin minha-feature
   ```
5. Abra um Pull Request.

---

## **Licença**
Este projeto está licenciado sob a Licença MIT. Veja o arquivo LICENSE para mais informações.
