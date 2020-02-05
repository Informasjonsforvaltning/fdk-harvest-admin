module.exports = {
  async up(db, client) {
    const session = client.startSession();

    try {
      await session.withTransaction(async () => {
        await db
          .collection('datasources')
          .updateMany(
            { dataSourceType: 'DCAT-AP-NO' },
            { $set: { dataType: 'dataset' } }
          );

        await db
          .collection('datasources')
          .updateMany(
            { dataSourceType: 'SKOS-AP-NO' },
            { $set: { dataType: 'concept' } }
          );
      });
    } finally {
      await session.endSession();
    }
  },

  async down(db) {
    await db
      .collection('datasources')
      .updateMany({}, { $set: { dataType: null } });
  }
};
