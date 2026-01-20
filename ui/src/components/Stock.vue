<template>
  <div>
    <h2>Query Stock</h2>
    <form @submit.prevent="queryStock">
      <div>
        <label for="stockId">Stock ID:</label>
        <input type="text" id="stockId" v-model="stockId">
      </div>
      <button type="submit">Query</button>
    </form>
  </div>
</template>

<script>
export default {
  data() {
    return {
      stockId: ''
    };
  },
  methods: {
    queryStock() {
      // Logic to send data to the backend
      console.log('Querying stock:', this.stockId);
    }
  }
};
</script>
