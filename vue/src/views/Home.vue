<template>
  <div>
    <div>
      <form @submit.prevent="startWithID">
        <input
          class="rounded-md mx-1 text-xl leading-none border-2 border-gray-300 w-3/4 py-2 tracking-widest text-center font-bold"
          pattern="^[0-9a-fA-F]{24}$" placeholder="TEST ID" ref="testID" required
          type="text"/>
        <button class="rounded-md mx-1 border-2 border-gray-300 py-3 w-1/6 bg-gray-500 text-white font-bold"
                type="submit">GO
        </button>
      </form>
    </div>


    <div class="mt-5">
      <button @click="startNew"
              class="rounded-md border-2 border-gray-300 py-3 w-11/12 bg-gray-500 text-white font-bold" id="new"
              name="new" type="submit"
              value="true">NEW TEST
      </button>
    </div>
  </div>
</template>

<script>
  export default {
    name: 'Home',
    methods: {
      startWithID() {
        this.$router.push({name: 'Load', params: {id: this.$refs.testID.value}})
      },
      startNew() {
        this.$router.push({name: 'New'})
      }
    }
  }
</script>

<style scoped>

</style>
