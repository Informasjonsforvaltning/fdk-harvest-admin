import mongoose from 'mongoose';

const DatasourceSchema = new mongoose.Schema(
  {
    id: {
      type: String,
      required: true,
      unique: true
    },
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
