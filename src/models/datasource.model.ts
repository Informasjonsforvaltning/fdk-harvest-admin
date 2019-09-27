import mongoose from 'mongoose';

const DatasourceSchema = new mongoose.Schema(
  {
    id: String,
    dataSourceType: String,
    url: String,
    publisher: String,
    description: String
  },
  {
    timestamps: true,
    strict: true
  }
);

export const DatasourceModel = mongoose.model('Datasource', DatasourceSchema);
