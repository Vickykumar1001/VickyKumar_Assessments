import { sendNotification } from '../src/a4_Notification_Service.js';
describe("Notification Service", function () {
    let mockNotificationService;

    beforeEach(function () {
        mockNotificationService = {
            send: jasmine.createSpy("send"),
        };
    });

    describe("sendNotification", function () {
        it("should return 'Notification Sent' when the notification is sent successfully", function () {
            mockNotificationService.send.and.returnValue(true);

            const message = "Hello World";
            const result = sendNotification(mockNotificationService, message);

            expect(mockNotificationService.send).toHaveBeenCalledWith(message);
            expect(result).toBe("Notification Sent");
        });

        it("should return 'Failed to Send' when the notification fails to send", function () {
            mockNotificationService.send.and.returnValue(false);

            const message = "Hello World";
            const result = sendNotification(mockNotificationService, message);

            expect(mockNotificationService.send).toHaveBeenCalledWith(message);
            expect(result).toBe("Failed to Send");
        });

        it("should not send a notification if no message is provided", function () {
            mockNotificationService.send.and.returnValue(false);

            const message = "";
            const result = sendNotification(mockNotificationService, message);

            expect(mockNotificationService.send).toHaveBeenCalledWith(message);
            expect(result).toBe("Failed to Send");
        });
    });
});
