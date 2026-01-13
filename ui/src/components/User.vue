<template>
  <div>
    <h2>Users</h2>
    <form @submit.prevent="createUser">
      <input v-model="newUser.name" placeholder="Name" required />
      <input v-model="newUser.email" placeholder="Email" required />
      <button type="submit">Create User</button>
    </form>
    <ul>
      <li v-for="user in users" :key="user.id">
        {{ user.name }} - {{ user.email }}
      </li>
    </ul>
  </div>
</template>

<script>
export default {
  data() {
    return {
      users: [],
      newUser: {
        name: '',
        email: ''
      }
    };
  },
  mounted() {
    this.fetchUsers();
  },
  methods: {
    async fetchUsers() {
      try {
        const response = await fetch('/api/v1/users');
        if (!response.ok) {
          throw new Error('Failed to fetch users');
        }
        const responseData = await response.json();
        this.users = responseData.data;
      } catch (error) {
        console.error(error);
      }
    },
    async createUser() {
      try {
        const response = await fetch('/api/v1/users', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json'
          },
          body: JSON.stringify(this.newUser)
        });
        if (!response.ok) {
          throw new Error('Failed to create user');
        }
        this.newUser.name = '';
        this.newUser.email = '';
        await this.fetchUsers();
      } catch (error) {
        console.error(error);
      }
    }
  }
};
</script>
