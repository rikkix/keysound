<template>
  <div>
    <ul>
      <li @click="judge(data)" v-for="data in getChoices()">
        <button class="rounded-md border-2 border-gray-300 py-3 w-11/12 bg-gray-500 text-white font-bold my-2">
          {{ data.name }}
        </button>
      </li>
    </ul>

  </div>
</template>

<script>
  function shuffle(array) {
    var currentIndex = array.length, randomIndex;

    // While there remain elements to shuffle...
    while (0 !== currentIndex) {
      // Pick a remaining element...
      randomIndex = Math.floor(Math.random() * currentIndex);
      currentIndex--;

      // And swap it with the current element.
      [array[currentIndex], array[randomIndex]] = [
        array[randomIndex], array[currentIndex]];
    }

    return array;
  }

  export default {
    name: "quiz",
    props: ['currentQuiz', 'nextQuiz', 'each'],
    methods: {
      getChoices() {
        var items = [];
        for (var i = 0; i < this.currentQuiz.choices.length; i++) {
          items.push({id: i, name: this.currentQuiz.choices[i], correct: false})
        }
        items.push({
          id: this.currentQuiz.choices.length,
          name: this.currentQuiz.name, correct: true
        });
        return shuffle(items)
      },
      judge(item) {
        if (item.correct) {
          return this.nextQuiz(this.each)
        } else {
          return this.nextQuiz(0)
        }
      }
    }
  }
</script>

<style scoped>

</style>
