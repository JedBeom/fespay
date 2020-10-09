export default function(record) {
    record.paidAt = new Date(record.paidAt);
    record.hours = record.paidAt.getHours();
    record.minutes = record.paidAt.getMinutes();
    record.seconds = record.paidAt.getSeconds();
    if (record.hours < 10 ) {
        record.hours = "0" + record.hours
    }
    if (record.minutes < 10 ) {
        record.minutes = "0" + record.minutes
    }
    if (record.seconds < 10 ) {
        record.seconds = "0" + record.seconds
    }

    if (record.canceledAt == undefined) {
        return record
    }

    record.canceledAt = new Date(record.canceledAt);
    record.canceledAt.hours = record.canceledAt.getHours();
    record.canceledAt.minutes = record.canceledAt.getMinutes();
    record.canceledAt.seconds = record.canceledAt.getSeconds();
    if (record.canceledAt.hours < 10 ) {
        record.canceledAt.hours = "0" + record.canceledAt.hours
    }
    if (record.canceledAt.minutes < 10 ) {
        record.canceledAt.minutes = "0" + record.canceledAt.minutes
    }
    if (record.canceledAt.seconds < 10 ) {
        record.canceledAt.seconds = "0" + record.canceledAt.seconds
    }

    return record
}