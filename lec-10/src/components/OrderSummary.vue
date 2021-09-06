<template>
  <div class="order-summary">
    <img :src="require('@/assets/images/illustration-hero.svg')" alt="illustration-hero.svg">
    <div class="content">

      <div class="content-text">
        <h1>Order Summary</h1>

        <p>You can now listen to millions of songs, audiobooks, and podcasts on any
          device anywhere you like!</p>
      </div>

      <div class="price">
        <div class="price-content">
          <div class="price-icon">
            <img :src="require('@/assets/images/icon-music.svg')" alt="icon-music.svg">
          </div>
          <div class="price-data">
            <h5>{{ this.package.title }}</h5>
            <p>{{ this.package.price }}/{{ this.package.duration }}</p>
          </div>
        </div>
        <p class="price-change" @click="changePackage">Change</p>
      </div>

      <div class="process">
        <button class="process-payment" @click="showModalItem"> Proceed to Payment</button>
        <p>Cancel Order</p>
      </div>
    </div>

    <ModalItem class="modal" v-show="showModal" @closeModalEvt="closeModalFunc($event)"/>
  </div>

</template>

<script>
import ModalItem from './ModalItem'

export default {
  components: {
    ModalItem
  },
  data() {
    return {
      package: {
        price: 59.99,
        title: "Annual Plan",
        duration: "year"
      },
      showModal: false
    }
  },
  methods: {
    changePackage() {
      if (this.package.title === "Annual Plan") {
        this.package.price = 99.99;
        this.package.title = "Pro Plan";
      } else {
        this.package.price = 59.99;
        this.package.title = "Annual Plan";
      }
    },
    showModalItem() {
      this.showModal = true;
    },
    closeModalFunc() {
      this.showModal = false;
    }
  }
}
</script>

<style scoped>
.order-summary {
  text-align: center;
  width: 25%;
  height: 80%;
  background: #ffffff;
  border-radius: 15px;
  margin: 5% auto;
}

.order-summary img {
  width: 100%;
  border-radius: 15px 15px 0 0;
}

.order-summary:hover {
  box-shadow: 0 8px 16px 0 rgba(0, 0, 0, 0.2);
}

.content {
  width: 80%;
  height: 100%;
  margin: 10% auto;
  text-align: center;
}

.content-text {
  display: inline-block;
  width: 80%;
}

.content-text h1 {
  font-size: 2.8rem;
  margin-bottom: 5%;
}

.price {
  display: flex;
  justify-content: space-around;
  align-items: center;
  background-color: hsl(225, 100%, 98%);
  border-radius: 15px;
  margin: 5% auto;
  height: 15%;
}

.price-change {
  color: hsl(245, 75%, 52%);
  text-decoration: underline;
}

.price-change:hover {
  cursor: pointer;
}

.price-data {
  width: 90%;
}

.price-data h5 {
  font-size: 1.6rem;
  font-weight: bold;
}

.price-icon {
  margin: auto 8% auto auto;
}

.price-content {
  display: flex;
}

.process {
  display: flex;
  flex-direction: column;
  height: 100%;
}

.process-payment {
  background-color: hsl(245, 75%, 52%);
  color: white;
  font-size: 1.6rem;
  font-weight: bold;
  border-radius: 10px;
  margin: 5% auto;
  width: 100%;
  height: 8%;
  cursor: pointer;
}

.process p {
  font-size: 1.6rem;
  font-weight: bold;
  cursor: pointer;
}

.modal {
  position: absolute;

}
</style>
