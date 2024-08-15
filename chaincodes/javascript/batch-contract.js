'use strict';

const { Contract } = require('fabric-contract-api');

class BatchContract extends Contract {

    async batchExists(ctx, batchId) {
        const buffer = await ctx.stub.getState(batchId);
        return (!!buffer && buffer.length > 0);
    }

    async userExists(ctx, userId) {
        const buffer = await ctx.stub.getState(userId);
        return (!!buffer && buffer.length > 0);
    }

    async createBatch(ctx, batchId, value) {

        console.log(batchId, value)

        const exists = await this.batchExists(ctx, batchId);
        if (exists) {
            throw new Error(`The batch ${batchId} already exists`);
        }

        const buffer = Buffer.from(JSON.stringify(value));
        await ctx.stub.putState(batchId, buffer);
    }

    async getBatch(ctx, batchId) {
        const buffer = await ctx.stub.getState(batchId);
        const exists = await this.batchExists(ctx, batchId);
        if (!exists) {
            throw new Error(`The batch ${batchId} does not exist`);
        }

        console.log("asset1", buffer.toString())

        const asset = JSON.parse(buffer.toString());

        console.log("asset", asset)

        return asset;
    }

    async getAllBatches(ctx) {

        const iterator = await ctx.stub.getStateByRange('', '');
        const allResults = [];

        while (true) {
            const res = await iterator.next();
            console.log("res.value:", res.value)
            if (res.value) {
                const strValue = Buffer.from(res.value.value).toString('utf8');
                console.log("strValue", strValue)
                let record;
                try {
                    record = JSON.parse(strValue);
                    console.log("record1", record)
                } catch (err) {
                    console.error(err);
                    record = strValue;
                }

                console.log("record", record)

                // console.log("record.value", typeof (record.value))
                var recordJSON = JSON.parse(record)
                console.log("recordJSON", recordJSON)

                if (recordJSON.docType == "Batch") {
                    allResults.push(recordJSON);
                }
            }

            console.log("allResults", allResults)

            if (res.done) {
                await iterator.close();
                return JSON.stringify(allResults);
            }
        }
    }

    async updateBatch(ctx, batchId, newValue) {
        const exists = await this.batchExists(ctx, batchId);
        if (!exists) {
            throw new Error(`The batch ${batchId} does not exist`);
        }

        const buffer = Buffer.from(JSON.stringify(newValue));
        await ctx.stub.putState(batchId, buffer);
    }

    async deleteBatch(ctx, batchId) {
        const exists = await this.batchExists(ctx, batchId);
        if (!exists) {
            throw new Error(`The batch ${batchId} does not exist`);
        }
        await ctx.stub.deleteState(batchId);
    }

    async createUser(ctx, userId, value) {

        console.log(userId, value)

        const exists = await this.userExists(ctx, userId);
        if (exists) {
            throw new Error(`The user ${userId} already exists`);
        }

        const buffer = Buffer.from(JSON.stringify(value));
        await ctx.stub.putState(userId, buffer);
    }


    async getUser(ctx, userId) {
        const buffer = await ctx.stub.getState(userId);
        const exists = await this.userExists(ctx, userId);
        if (!exists) {
            throw new Error(`The user ${userId} does not exist`);
        }

        console.log("user1", buffer.toString())

        const user = JSON.parse(buffer.toString());

        console.log("user", user)

        return user;
    }
}

module.exports = BatchContract;
