<template>
  <div>
    <p>Requesting...</p>
    <FlashMessage :position="'right bottom'"/>
  </div>
</template>

<script>
  import axios from 'axios'

  export default {
    name: 'New',
    mounted() {
      axios.post("/api/new").then((res) => {
        this.$router.push({name: 'Load', params: {id: res.data.id, quizInfo: res.data}})
      }).catch((err) => {
        console.log(err);
        this.flashMessage.error({
          title: err.name + " (" + err.config.url + ")",
          message: err.message,
        });
      })
    }
  }
</script>

<style scoped>

</style>
