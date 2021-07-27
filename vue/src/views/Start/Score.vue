<template>
  <div>
    <h1 class="text-6xl tracking-wide	font-extrabold text-gray-900 mt-5">
      {{ quizScore }}/100</h1>
    <div @click="copyLink" class="mt-3 font-bold cursor-pointer">
      TEST ID: {{ $route.params.id }}
    </div>
    <div class="mt-5">
      <div @click="showAnswer=!showAnswer" class="cursor-pointer">
        <p class="underline" v-if="!showAnswer">show correct answer</p>
        <p class="underline" v-else>hide correct answer</p>
      </div>
      <table class="border-collapse border border-gray-900 w-full mt-2" v-show="showAnswer">
        <tr v-for="(data,index) in quizInfo.pieces">
          <td class="w-1/12  border border-gray-900">
            {{ index+1 }}
          </td>
          <td class="w-11/12 border border-gray-900">
            {{ data.title }}
          </td>
        </tr>
      </table>
    </div>
  </div>
</template>

<script>
  export default {
    name: "Score",
    props: ['quizScore', 'quizInfo'],
    data() {
      return {
        showAnswer: false,
        shareLink: window.location.protocol + "//" + window.location.host + "/start/" + this.$route.params.id,
      }
    },
    methods: {
      copyLink() {
        window.prompt("Copy to clipboard: Ctrl+C, Enter", this.shareLink);
      }
    },
    mounted() {
      console.log(this.$route)
    }
  }
</script>

<style scoped>

</style>
