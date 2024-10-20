<template>
  <div>
    <el-header>
      <div class="logo">Medilearn</div>
      <el-menu mode="horizontal" class="menu">
        <el-menu-item index="1">Accueil</el-menu-item>
        <el-menu-item index="2">Cours</el-menu-item>
        <el-menu-item index="3">À propos</el-menu-item>
        <el-menu-item index="4">Se connecter</el-menu-item>
      </el-menu>
    </el-header>

    <el-container>
      <el-main>
        <div class="hero">
          <h1>Bienvenue sur Medilearn</h1>
          <p>Votre plateforme pour télécharger et partager des cours.</p>
          <el-button type="primary">Voir les cours</el-button>
          <el-button type="success">Rejoignez-nous</el-button>
        </div>

        <el-row class="features" type="flex" justify="center">
          <el-col v-for="(feature, index) in features" :key="index" :span="6">
            <el-card :body-style="{ padding: '20px' }">
              <h3>{{ feature.title }}</h3>
              <p>{{ feature.description }}</p>
            </el-card>
          </el-col>
        </el-row>
      </el-main>
    </el-container>

    <el-footer>
      <p>© 2024 Medilearn. Tous droits réservés.</p>
    </el-footer>
  </div>
</template>

<script>
export default {
  data() {
    return {
      features: [
        { title: 'Télécharger des cours', description: 'Accédez à une variété de cours disponibles.' },
        { title: 'Partager vos cours', description: 'Contribuez à la communauté en téléchargeant vos propres cours.' },
        { title: 'Feedback de la communauté', description: 'Laissez des commentaires et des évaluations sur les cours.' },
        { title: 'Gestion par les modérateurs', description: 'Les cours téléchargés sont vérifiés par nos modérateurs.' },
      ],
    };
  },
};
</script>

<style scoped>
.logo {
  font-size: 24px;
  font-weight: bold;
}

.hero {
  text-align: center;
  padding: 50px 0;
}

.features {
  margin-top: 20px;
}

.el-footer {
  text-align: center;
  padding: 20px;
  background-color: #f1f1f1;
}
</style>

